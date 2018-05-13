package responses

import "net/http"

type location string

func (l location) Write(w http.ResponseWriter) {
	w.Header().Set("Location", string(l))
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Location")
	w.WriteHeader(302)
}

func Created(l string) Response {
	return location(l)
}
