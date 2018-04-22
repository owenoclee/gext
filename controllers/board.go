package controllers

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/datastore"
	"github.com/owenoclee/gext-server/responses"
)

var ShowBoard Action = func(r *http.Request, p httprouter.Params, ds datastore.Datastore) responses.Response {
	// Validate parameters
	pageNum, err := strconv.ParseInt(p.ByName("page"), 10, 64)
	if err != nil {
		return responses.Status(422)
	}
	if pageNum < 1 {
		pageNum = 1
	}

	page, err := ds.GetPage(p.ByName("board"), pageNum)
	if err != nil {
		return responses.LogError(err)
	} else if page.GetThreads() == nil {
		return responses.Status(404)
	}
	return responses.Protobuf(&page, 200)
}
