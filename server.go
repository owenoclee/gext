package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

type serverConfig struct {
	Address string
}

func serveHTTP(router *httprouter.Router) {
	var config serverConfig
	loadConfig(&config, "cfg/server.json")
	log.Fatal(http.ListenAndServe(config.Address, router))
}
