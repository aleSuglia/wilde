package main

import (
	"testing"
)

var (
	termsFilename = "terms.txt"
)

func TestTokenizer(t *testing.T) {
	terms, err := TokenizeTextFile(termsFilename, ',')

	if err != nil {
		t.Fatalf("Unable to tokenize file: %+v", err)
	}

	for _, term := range terms {
		t.Logf("%+v", term)
	}

}
