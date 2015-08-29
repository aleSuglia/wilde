package main

import (
	"testing"
)

var (
	terms = []string{"dog", "milk", "dumb"}
)

func init() {
	LoadConfiguration("wilde.json")

	if err := OpenConnection(Configuration.Host, Configuration.DbName); err != nil {
		panic(err)
	}

}

func TestGetTerms(t *testing.T) {
	t.Logf("Translating %+v", terms)

	if _, err := GetAllTerms(terms, Configuration.FromLang, Configuration.ToLang); err != nil {
		t.Fatalf("Unable to insert all terms %+v", err)
	}

	t.Log("Correctly inserted all the terms")
}

func TestInsertTerms(t *testing.T) {
	t.Logf("Translating %+v", terms)

	transMap, err := GetAllTerms(terms, Configuration.FromLang, Configuration.ToLang)
	if err != nil {
		t.Fatalf("(get_all_terms): %+v", err)
	}

	t.Log("Correctly inserted all the terms")

	// select the first translation for the first term
	selMap := map[string]WRData{
		terms[0]: transMap[terms[0]][0],
	}

	if err := InsertTermsTranslations(selMap); err != nil {
		t.Fatalf("(insert_terms_transl): %+v", err)
	}

	t.Log("Correctly inserted")

}
