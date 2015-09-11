package main

import (
	"fmt"

	"gopkg.in/mgo.v2/bson"
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

	var termData Term

	if err := termsCol.Find(bson.M{"token": term}).One(&termData); err != nil {
		fmt.Printf("(GetTranslations - Term): %+v\n", err)
		return nil, err
	}

	transData := make([]TermTrans, len(termData.Trans))

	for i, trans := range termData.Trans {
		if err := termsTransCol.FindId(trans).One(&transData[i]); err != nil {
			fmt.Printf("(GetTranslations - TermTrans): %+v\n", err)
			return nil, err
		}
	}

	return transData, nil
}

func dbDump() {
	termsCol := DBObj.C(termsCollName)

	num, err := termsCol.Count()
	if err != nil {
		fmt.Printf("Count failed on termsCol:  %q\n", err)
		return
	}
	fmt.Println("TermsCol count", num)

	iter := termsCol.Find(nil).Iter()
	var t Term
	for iter.Next(&t) {
		fmt.Printf("%#v tnum: %d\n", t, len(t.Trans))
	}

	termsTransCol := DBObj.C(termsTransCollName)
	num, err = termsTransCol.Count()
	if err != nil {
		fmt.Printf("Count failed on termsTransCol:  %q\n", err)
		return
	}
	fmt.Println("TermsTransCol count", num)

	iter = termsTransCol.Find(nil).Iter()
	tt := TermTrans{}
	for iter.Next(&tt) {
		fmt.Printf("%#v\n", tt)
	}
	fmt.Printf("\n\n\n")
}
