package main

import (
	"encoding/json"
	"os"
)

func loadConfig(config interface{}, filePath string) {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	err = decoder.Decode(config)
	if err != nil {
		panic(err)
	}
}
