package controllers

import (
	"html/template"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/responses"
)

type Action func(*http.Request, httprouter.Params, datastore.Datastore, *template.Template) responses.Response

func (a Action) Handler(ds datastore.Datastore, t *template.Template) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		response := a(r, p, ds, t)
		response.Write(w)
	}
}
