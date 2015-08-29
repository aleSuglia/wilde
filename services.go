package main

import (
	"errors"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"net/http"
)

const (
	//WordReference link for translation
	// In order to get all the translation for the word "make" from english
	// to italian you need to go to use as URL path: "/enit/make"
	wordReferenceLinkFormat = "http://wordreference.com/%s%s/%s"
	FromTransTag            = "FrWrd"
	ToTransTag              = "ToWrd"
	EmptyTag                = "wrtopsection"
	MainTranslationTable    = "WRD"
	FromLangExample         = "FrEx"
	ToLangExample           = "ToEx"
)

var (
	ErrTranslationPage = errors.New("Unable to get translation page for term")
	ErrParsePage       = errors.New("Unable to parse the translation page")
)

type WRData struct {
	TermUse         string
	OtherInfo       string
	Trans           string
	OrigLangExample string
	DstLangExample  string
	FromLang        string
	ToLang          string
}

func GetTranslationPage(term string, fromLang, toLang string) ([]WRData, error) {
	resp, err := http.Get(fmt.Sprintf(wordReferenceLinkFormat, fromLang, toLang, term))
	if err != nil {
		return nil, ErrTranslationPage
	}

	defer resp.Body.Close()
	if trans, err := parsePage(resp.Body, fromLang, toLang); err == nil {
		return trans, nil
	}

	return nil, ErrParsePage
}

func hasAttributeValue(n *html.Node, key, value string) bool {
	for _, val := range n.Attr {
		if val.Key == key && val.Val == value {
			return true
		}
	}

	return false
}

func hasAttribute(n *html.Node, key string) (string, bool) {
	for _, val := range n.Attr {
		if val.Key == key {
			return val.Val, true
		}
	}

	return "", false
}

func retrieveTranslationTables(n *html.Node, tables *[]*html.Node) {
	if n.Type == html.ElementNode && hasAttributeValue(n, "class", MainTranslationTable) {
		*tables = append(*tables, n.FirstChild)
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			retrieveTranslationTables(c, tables)
		}
	}

}

func parsePage(body io.Reader, fromLang, toLang string) ([]WRData, error) {
	doc, err := html.Parse(body)
	if err != nil {
		return nil, ErrParsePage
	}

	var transTables []*html.Node

	retrieveTranslationTables(doc, &transTables)

	// Here we got the table, if not, the pointer should be null
	if len(transTables) != 0 {
		var translations []WRData

		for _, table := range transTables {
			table, err := parseTable(table, fromLang, toLang)
			if err != nil {
				return nil, ErrParsePage
			}

			translations = append(translations, table...)
		}

		return translations, nil
	}
	return nil, ErrParsePage
}

func parseTable(transTable *html.Node, fromLang, toLang string) ([]WRData, error) {
	var dataTrans []WRData

	// at this moment, there are tr elements
	// We should avoid wrtopsection;
	currSec := "even"
	currSecData := WRData{}
	for c := transTable.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.ElementNode {
			class, _ := hasAttribute(c, "class")

			if class == EmptyTag {
				// skip this tag anyway
				continue
			} else { // tr with class "even" or "odd"

				// if the current tr is not the same as before
				// add the collected information to the map
				if currSec != EmptyTag {
					if class != currSec {
						currSecData.FromLang = fromLang
						currSecData.ToLang = toLang
						dataTrans = append(dataTrans, currSecData)
						currSecData = WRData{}
						currSec = class
					}
				} else {
					// restore it
					currSec = class
				}

				// ##assert##: same class value

				// wrtopsection in an odd or even tr
				if c.FirstChild != nil && hasAttributeValue(c.FirstChild, "class", EmptyTag) {
					currSec = EmptyTag
					continue
				} else if _, ok := hasAttribute(c, "id"); ok {
					//translation tr
					getTranslation(c, &currSecData)
				} else {
					// examples tr
					getExamples(c, &currSecData)
				}

			}
		}

	}
	return dataTrans, nil
}

func getExamples(c *html.Node, currSecData *WRData) {
	for innerC := c.FirstChild; innerC != nil; innerC = innerC.NextSibling {
		class, _ := hasAttribute(innerC, "class")

		switch class {
		case FromLangExample:
			currSecData.OrigLangExample = innerC.FirstChild.Data
			break
		case ToLangExample:
			currSecData.DstLangExample = innerC.FirstChild.Data
			break
		}
	}
}

func getTranslation(father *html.Node, currSecData *WRData) {
	for c := father.FirstChild; c != nil; c = c.NextSibling {
		class, _ := hasAttribute(c, "class")

		switch class {
		case EmptyTag:
			// NOP
			break
		case FromTransTag:
			currSecData.TermUse = c.FirstChild.FirstChild.Data
			break
		case ToTransTag:
			currSecData.Trans = c.FirstChild.Data
			break
		default:
			// other information about the term
			// could be not present!
			if c.FirstChild != nil {
				currSecData.OtherInfo = c.FirstChild.Data
			}
			break
		}
	}

}
