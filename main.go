package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
	"text/scanner"
	"time"
)

var (
	termsList      string
	period         string
	topN           int
	configFile     string
	transInfoTempl = "\n%d)Term-info:%s\t%s\nTranslation:%s\n"
)

func init() {
	flag.StringVar(&termsList, "says", "", "Stores information about new terms. Accept a string with terms delimited by comma or a single term")
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
	if err != nil || index > int64(len(termData))-1 {
		return -1
	}

	return index
}

func displayTranslations(data []WRData) {
	for i, t := range data {
		fmt.Printf(transInfoTempl, i, t.TermUse, t.OtherInfo, t.Trans)

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
			fmt.Println("Skipped translation")
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
		fmt.Println("Not enough parameters specified\n")
		flag.Usage()
		return
	}

	if f := flag.Lookup("config"); f == nil {
		fmt.Println("Configuration file not specified\n")
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
			if termsList == "" {
				return
			}

			dataTrans, err := GetAllTerms(arrayTerms(strings.Split(termsList, ",")), Configuration.FromLang, Configuration.ToLang)

			if err != nil {
				fmt.Printf("%+v", err)
				return
			}

			selData := getSelectedTranslations(dataTrans)

			if err := InsertTermsTranslations(selData); err != nil {
				panic(err)
			}

		case "top":
			visualizeTop(GetTopTerms(topN))

		case "stats":
			stats := GetTermStatistics()
			fmt.Printf("%s\n", stats)

		case "trans":
			//trans := GetTranslations(flag.Value)
			//fmt.Printf("%+v", trans)

		case "config":
			// NOP

		default:
			panic("Undefined operation")

		}

	})

}
