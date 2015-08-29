package main

import (
	"fmt"
)

var (
	transInfoTemplate = "%d)\nTerm-info:%s\t%s\nTranslation:%s\n"
)

func GetAllTerms(terms arrayTerms, langFrom, langTo string) (map[string][]WRData, error) {
	transData := make(map[string][]WRData)
	// for each term, download its page and get all the information
	for _, t := range terms {
		data, err := GetTranslationPage(t, langFrom, langTo)
		if err != nil {
			fmt.Printf("Failed get translation page: %+v", err)
			return nil, err
		}

		transData[t] = data
	}

	return transData, nil
}

func InsertTermsTranslations(termsTrans map[string]WRData) error {
	for t, data := range termsTrans {
		// Insert the term with the selected translation information
		if err := InsertTerm(t, data); err != nil {
			fmt.Printf("Failed insert term %s: %+v", t, err)
			return err
		}

		fmt.Printf("Correctly inserted <%s>", t)

	}

	return nil

}
