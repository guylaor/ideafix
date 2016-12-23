package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

var config map[string]interface{}

func loadConfigValues() {
	content, err := ioutil.ReadFile("conf.json")
	checkError(err)

	if err := json.Unmarshal(content, &config); err != nil {
		checkError(err)
	}

	log.Printf("Loaded config \n")

}
