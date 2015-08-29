package main

import (
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"time"
)

// Term defines the structure of a specific word of the language
// For each token, written in a given language, could be present
// different translations
type Term struct {
	_id   bson.ObjectId
	Token string
	Lang  string
	Trans []bson.ObjectId
}

// TermTrans defines a specific translation for a given term
type TermTrans struct {
	_id         bson.ObjectId
	TermID      bson.ObjectId
	Trans       string
	OtherInfo   string
	TermUse     string
	Lang        string
	Period      time.Time
	FromExample string
	ToExample   string
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
