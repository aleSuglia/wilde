package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"log"
	"time"
)

const (
	termsCollName      = "terms"
	termsTransCollName = "terms_trans"
)

var (
	DBObj *mgo.Database
)

func OpenConnection(host, dbName string) error {

	session, err := mgo.Dial(host)
	if err != nil {
		log.Printf("(OpenConnection) %+v", err)
		return err
	}

	//TODO: no password for database?
	DBObj = session.DB(dbName)

	return nil
}

func InsertTerm(term string, termsData WRData) error {

	termsCol := DBObj.C(termsCollName)
	termsTransCol := DBObj.C(termsTransCollName)

	termData := Term{}
	// check if already exists the term
	if err := termsCol.Find(Term{
		Token: term,
	}).One(&termData); err != nil {
		// term doesn't exist
		termData.Lang = termsData.FromLang
		termData.Token = term
		termData._id = bson.NewObjectId()

		termTransData := TermTrans{
			_id:         bson.NewObjectId(),
			TermID:      termData._id,
			TermUse:     termsData.TermUse,
			OtherInfo:   termsData.OtherInfo,
			Trans:       termsData.Trans,
			Lang:        termsData.ToLang,
			Period:      time.Now(),
			FromExample: termsData.OrigLangExample,
			ToExample:   termsData.DstLangExample,
		}

		termData.Trans = make([]bson.ObjectId, 0)
		termData.Trans = append(termData.Trans, termTransData._id)

		if err := termsCol.Insert(termData); err != nil {
			fmt.Printf("(InsertTerm): %+v", err)
			return err
		}

		if err := termsTransCol.Insert(termTransData); err != nil {
			fmt.Printf("(InsertTermTrans): %+v", err)
			return err
		}

	} else {
		// term already exists
		termTransData := TermTrans{
			_id:         bson.NewObjectId(),
			TermID:      termData._id,
			TermUse:     termsData.TermUse,
			OtherInfo:   termsData.OtherInfo,
			Trans:       termsData.Trans,
			Lang:        termsData.ToLang,
			Period:      time.Now(),
			FromExample: termsData.OrigLangExample,
			ToExample:   termsData.DstLangExample,
		}

		if err := termsTransCol.Insert(termTransData); err != nil {
			fmt.Printf("(InsertTermTrans): %+v", err)
			return err
		}

		termData.Trans = append(termData.Trans, termTransData._id)

		if err := termsCol.Update(Term{
			_id: termData._id,
		}, termData); err != nil {
			fmt.Printf("(UpdateTerm): %+v", err)
			return err
		}

	}

	return nil
}
