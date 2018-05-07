package responses

import (
	"fmt"
	"log"
	"net/http"
)

type serverError struct {
	log     bool
	message string
}

func (e serverError) Write(w http.ResponseWriter) {
	if e.log {
		log.Print(e.message)
	}
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(500)
}

func LogError(err interface{}) Response {
	return serverError{log: true, message: fmt.Sprint(err)}
}
