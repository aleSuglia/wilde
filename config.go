package main

import (
	"encoding/json"
	"io/ioutil"
	"os"
)

type Config struct {
	Host     string
	DbName   string
	FromLang string
	ToLang   string
}

var (
	Configuration Config
)

func LoadConfiguration(configurationFile string) {
	file, err := os.Open(configurationFile)

	if err != nil {
		panic(err)
	}

	buf, err := ioutil.ReadAll(file)

	if err != nil {
		panic(err)
	}

	if err := json.Unmarshal(buf, &Configuration); err != nil {
		panic(err)
	}
}
