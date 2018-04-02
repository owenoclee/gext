package responses

import "net/http"

type status int

func (s status) Write(w http.ResponseWriter) {
	w.WriteHeader(int(s))
}

func Status(s int) Response {
	return status(s)
}
