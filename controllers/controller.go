package controllers

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/datastore"
	"github.com/owenoclee/gext-server/responses"
)

type Action func(*http.Request, httprouter.Params, datastore.Datastore) responses.Response

func (a Action) Handler(ds datastore.Datastore) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		response := a(r, p, ds)
		response.Write(w)
	}
}
