package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

//GetAllTerms retrieves translations for each term in the list "terms"
//translation options are mapped to the lookup term
func GetAllTerms(terms arrayTerms, langFrom, langTo string) (map[string][]WRData, error) {
	transData := make(map[string][]WRData)
	// for each term, download its page and get all the information

	for _, t := range terms {
		data, err := GetTranslationPage(t, langFrom, langTo)
		if err != nil {
			fmt.Printf("Failed get translation page for term %q: %+v\n", t, err)
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

func InsertTerm(term string, termsData WRData) error {

	termsCol := DBObj.C(termsCollName)
	termsTransCol := DBObj.C(termsTransCollName)

	var tData Term

	// check if the term already exists
	if err := termsCol.Find(bson.M{"token": term}).One(&tData); err != nil {
		// term doesn't exist

		fmt.Println("Find Error: ", err)

		tData.Lang = termsData.FromLang
		tData.Token = term
		tData.Id = bson.NewObjectId()
		tData.Trans = make([]bson.ObjectId, 0)

		termTransData := TermTrans{
			Id:          bson.NewObjectId(),
			TermID:      tData.Id,
			TermUse:     termsData.TermUse,
			OtherInfo:   termsData.OtherInfo,
			Trans:       termsData.Trans,
			Lang:        termsData.ToLang,
			Period:      time.Now(),
			FromExample: termsData.OrigLangExample,
			ToExample:   termsData.DstLangExample,
		}

		tData.Trans = append(tData.Trans, termTransData.Id)

		fmt.Printf("Created Term: %#v\n", tData)

		if err := termsCol.Insert(tData); err != nil {
			fmt.Printf("(InsertTerm): %+v\n", err)
			return err
		}

		if err := termsTransCol.Insert(termTransData); err != nil {
			fmt.Printf("(InsertTermTrans): %+v\n", err)
			return err
		}

	} else {
		// term already exists
		termTransData := TermTrans{
			Id:          bson.NewObjectId(),
			TermID:      tData.Id,
			TermUse:     termsData.TermUse,
			OtherInfo:   termsData.OtherInfo,
			Trans:       termsData.Trans,
			Lang:        termsData.ToLang,
			Period:      time.Now(),
			FromExample: termsData.OrigLangExample,
			ToExample:   termsData.DstLangExample,
		}

		tData.Trans = append(tData.Trans, termTransData.Id)

		if err := termsTransCol.Insert(termTransData); err != nil {
			fmt.Printf("(InsertTermTrans): %+v\n", err)
			return err
		}

		if err := termsCol.Update(bson.M{"token": term}, tData); err != nil {
			fmt.Printf("(UpdateTerm): %+v\n", err)
			return err
		}

	}

	return nil
}
