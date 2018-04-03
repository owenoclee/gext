package main

import (
	"encoding/json"
	"log"
	"os"
)

func loadConfig(config interface{}, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	decoder.DisallowUnknownFields()
	if err := decoder.Decode(config); err != nil {
		log.Fatal(err)
	}
}
