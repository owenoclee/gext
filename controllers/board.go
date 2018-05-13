package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/responses"
)

var ShowBoard Action = func(_ *http.Request, p httprouter.Params, ds datastore.Datastore, t *template.Template) responses.Response {
	// Read
	pageNum64, err := strconv.ParseUint(p.ByName("page"), 10, 32)

	// Validate
	if err != nil {
		return responses.Status(422)
	}
	pageNum := uint32(pageNum64)
	if pageNum < 1 {
		pageNum = 1
	}

	// Show
	board := p.ByName("board")
	page, err := ds.GetPage(board, pageNum)
	if err != nil {
		return responses.LogError(err)
	} else if len(page.Threads) == 0 {
		return responses.Status(404)
	}
	return responses.View(
		t.Lookup("board.html"),
		responses.ViewData{
			Title: fmt.Sprintf("/%v/ - gext", board),
			Data:  page,
		},
	)
}
