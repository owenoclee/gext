package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"

	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/models"
	"github.com/owenoclee/gext/responses"
	"goji.io/pat"
	"goji.io/pattern"
)

var ShowBoard Action = func(r *http.Request, ds datastore.Datastore, t *template.Template) responses.Response {
	// Read
	var pageNum64 uint64
	var err error
	if _, ok := r.Context().Value(pattern.Variable("page")).(string); ok {
		pageNum64, err = strconv.ParseUint(pat.Param(r, "page"), 10, 32)
	} else {
		pageNum64 = 1
	}

	// Validate
	if err != nil {
		return responses.Status(422)
	}
	pageNum := uint32(pageNum64)
	if pageNum < 1 {
		pageNum = 1
	}

	// Show
	board := pat.Param(r, "board")
	page, err := ds.GetPage(board, pageNum)
	if err != nil {
		return responses.LogError(err)
	}
	return responses.View(
		t.Lookup("board.html"),
		responses.ViewData{
			Title: fmt.Sprintf("/%v/ - gext", board),
			Data: struct {
				Page  models.Page
				Board string
			}{
				page,
				board,
			},
		},
	)
}
