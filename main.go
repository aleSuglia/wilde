package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"text/scanner"
	"time"
)

var (
	termsList  arrayTerms
	period     string
	topN       int
	configFile string
)

func init() {
	flag.Var(&termsList, "says", "Stores information about new terms. Accept a string with terms delimited by comma or a single term")
	flag.IntVar(&topN, "top", 10, "Return topN most frequently used terms - default top 10")
	flag.StringVar(&period, "stats", time.Now().Format(time.UnixDate), "Return statistics about terms in a specified period")
	flag.String("trans", "", "Get all the translations for a specific keyword ordered by frequency")
	flag.StringVar(&configFile, "config", "", "Insert the path of the configuration file that wilde will use")

}

func selectTranslation(termData []WRData) int64 {
	scan := scanner.Scanner{}
	scan.Init(os.Stdin)
	fmt.Print("Select one of the translations (e.g., 1): ")
	scan.Scan()

	if !scan.IsValid() {
		return -1
	}

	index, err := strconv.ParseInt(scan.TokenText(), 10, 64)
	if err != nil || index > int64(len(termData)) {
		return -1

	}

	return index
}

func displayTranslations(data []WRData) {
	for i, t := range data {
		fmt.Printf(transInfoTemplate, i, t.TermUse, t.OtherInfo, t.Trans)

		if t.OrigLangExample != "" {
			fmt.Printf("%s example: %s\n\n", t.FromLang, t.OrigLangExample)
		}

		if t.DstLangExample != "" {
			fmt.Printf("%s example: %s\n\n", t.ToLang, t.DstLangExample)
		}
	}

}

func getSelectedTranslations(dataTrans map[string][]WRData) map[string]WRData {
	selTransData := make(map[string]WRData)

	fmt.Printf("Translation from \"%s\" to \"%s\"\n", Configuration.FromLang, Configuration.ToLang)
	for t, data := range dataTrans {
		displayTranslations(data)
		transIndex := selectTranslation(data)

		if transIndex == -1 {
			fmt.Printf("Skipped translation")
			continue
		}

		selTransData[t] = data[transIndex]
	}

	return selTransData
}

func visualizeTop(tops []string) {}

func main() {
	flag.Parse()

	// Not enough parameters
	if flag.NFlag() != 2 {
		fmt.Printf("Not enough parameters specified\n\n")
		flag.Usage()
		return
	}

	if f := flag.Lookup("config"); f == nil {
		fmt.Printf("Configuration file not specified\n\n")
		flag.Usage()
		return
	}

	LoadConfiguration(configFile)

	// initialize database connection
	if err := OpenConnection(Configuration.Host, Configuration.DbName); err != nil {
		panic(err)
	}

	flag.Visit(func(flag *flag.Flag) {
		switch flag.Name {
		case "says":
			dataTrans, err := GetAllTerms(termsList, Configuration.FromLang, Configuration.ToLang)

			if err != nil {
				fmt.Printf("%+v", err)
				return
			}

			selData := getSelectedTranslations(dataTrans)

			if err := InsertTermsTranslations(selData); err != nil {
				panic(err)
			}
			break
		case "top":
			visualizeTop(GetTopTerms(topN))
			break
		case "stats":
			stats := GetTermStatistics()
			fmt.Printf("%s", stats)
			break
		case "trans":
			//trans := GetTranslations(flag.Value)
			//fmt.Printf("%+v", trans)
			break
		case "config":
			// NOP
			break
		default:
			panic("Undefined operation")

		}

	})

}
