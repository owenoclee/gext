package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"strings"

	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/models"
	"github.com/owenoclee/gext/responses"
	"goji.io/pat"
)

var StorePost Action = func(r *http.Request, ds datastore.Datastore, t *template.Template) responses.Response {
	// Read
	r.ParseForm()
	replyTo, err := strconv.ParseUint(pat.Param(r, "id"), 10, 32)
	if err != nil {
		return responses.Status(400)
	}
	post := models.Post{
		ReplyTo: uint32(replyTo),
		Body:    r.FormValue("body"),
	}

	// Validate
	post.Body = strings.TrimSpace(post.Body)
	if post.Body == "" || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}
	board, err := ds.GetThreadBoard(post.ReplyTo)
	if board == "" || board != pat.Param(r, "board") {
		if err != nil {
			return responses.LogError(err)
		}
		return responses.Status(422)
	}

	// Store
	id, err := ds.StorePost(post)
	if err != nil {
		return responses.LogError(err)
	}
	return responses.Created(fmt.Sprintf("/%v/thread/%v#%v", board, post.ReplyTo, id))
}
