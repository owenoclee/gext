package controllers

import (
	"database/sql"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/responses"
)

type Action func(*http.Request, httprouter.Params, *sql.DB) responses.Response

func (a Action) Handler(db *sql.DB) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		response := a(r, p, db)
		response.Write(w)
	}
}
