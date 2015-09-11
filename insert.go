package main

import (
	"fmt"
)

func GetAllTerms(terms arrayTerms, langFrom, langTo string) (map[string][]WRData, error) {
	transData := make(map[string][]WRData)
	// for each term, download its page and get all the information
	fmt.Println(terms[0])
	for _, t := range terms {
		data, err := GetTranslationPage(t, langFrom, langTo)
		if err != nil {
			fmt.Printf("Failed get translation page: %+v\n", err)
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
			fmt.Printf("Failed insert term %s: %+v\n", t, err)
			return err
		}

		fmt.Printf("Correctly inserted <%s>\n", t)

	}

	return nil

}
