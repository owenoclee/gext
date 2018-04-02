package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func serveHTTP(router *httprouter.Router) {
	log.Fatal(http.ListenAndServe(":8080", router))
}
