package controllers

import (
	"html/template"
	"net/http"

	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/responses"
)

type Action func(*http.Request, datastore.Datastore, *template.Template) responses.Response

func (a Action) Handler(ds datastore.Datastore, t *template.Template) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		response := a(r, ds, t)
		response.Write(w)
	})
}
