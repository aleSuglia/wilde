package main

import (
	"fmt"
)

func GetTopTerms(topN int) (arr []string) {
	return
}

func GetTermStatistics() (str string) {
	return
}

func GetTranslations(term string) ([]TermTrans, error) {
	termsCol := DBObj.C(termsCollName)
	termsTransCol := DBObj.C(termsTransCollName)

	termData := Term{}

	if err := termsCol.Find(Term{
		Token: term,
	}).One(&termData); err != nil {
		fmt.Printf("(GetTranslations): %+v", err)
		return nil, err
	}

	transData := make([]TermTrans, len(termData.Trans))

	for i, trans := range termData.Trans {
		if err := termsTransCol.Find(TermTrans{
			_id: trans,
		}).One(&transData[i]); err != nil {
			fmt.Printf("(GetTranslations): %+v", err)
			return nil, err
		}
	}

	return transData, nil
}
