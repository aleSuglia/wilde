package main

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Term defines the structure of a specific word of the language
// For each token, written in a given language, could be present
// different translations
type Term struct {
	Id    bson.ObjectId   `bson:"_id,omitempty"`
	Token string          `bson:"token"`
	Lang  string          `bson:"lang"`
	Trans []bson.ObjectId `bson:"trans"`
}

// TermTrans defines a specific translation for a given term
type TermTrans struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`
	TermID      bson.ObjectId `bson:"termid"`
	Trans       string        `bson:"trans"`
	OtherInfo   string        `bson:"otherinfo"`
	TermUse     string        `bson:"termuse"`
	Lang        string        `bson:"lang"`
	Period      time.Time     `bson:"period"`
	FromExample string        `bson:"fromexample"`
	ToExample   string        `bson:"toexample"`
}

type arrayTerms []string

func (arr *arrayTerms) String() string {
	str := ""
	for i, t := range []string(*arr) {
		str += fmt.Sprintf("%d) %s\n", i, t)
	}

	return str
}

func (arr *arrayTerms) Set(value string) error {
	*arr = append(*arr, value)
	return nil
}

func (t *TermTrans) String() (s string) {
	s = fmt.Sprintf("ID: %v \t termID: %v \n Trans: %s \n Other Info: %s \n", t.Id, t.TermID, t.Trans, t.OtherInfo)
	s += fmt.Sprintf(" Term Use: %s \n Language: %s \n Period: %v \n", t.TermUse, t.Lang, t.Period)
	s += fmt.Sprintf(" Native Example: %s \n Translation Example: %s\n\n", t.FromExample, t.ToExample)
	return
}
