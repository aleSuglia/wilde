package main

import (
	"log"

	"gopkg.in/mgo.v2"
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
		log.Printf("(OpenConnection) %+v\n", err)
		return err
	}

	session.SetSafe(&mgo.Safe{})

	//TODO: no password for database?
	DBObj = session.DB(dbName)

	return nil
}
