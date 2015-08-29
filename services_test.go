package main

import (
	_ "os"
	"testing"
)

const (
	FromLang = "en"
	ToLang   = "it"
	term     = "chill"
	testPage = "test_trans.html"
)

func TestGetPage(t *testing.T) {
	t.Logf("Translation page for the term %s", term)

	trans, err := GetTranslationPage(term, FromLang, ToLang)

	if err != nil {
		t.Error(err)
		t.Fail()
	}

	t.Log("Correctly parsed page")
	t.Logf("%+v", trans)
}

/*func TestParsePage(t *testing.T) {
	t.Logf("Translation page for the term %s", term)

	file, err := os.Open(testPage)

	if err != nil {
		t.Fatal(err)
	}

	defer file.Close()

	trans, err := parsePage(file)

	if err != nil {
		t.Fatal(err)
	}

	t.Log("Correctly parsed page")

	for i, val := range trans {
		t.Logf("%d) %+v", i, val)
	}
}*/
