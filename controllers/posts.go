package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/models"
	"github.com/owenoclee/gext/responses"
)

var StorePost Action = func(r *http.Request, _ httprouter.Params, ds datastore.Datastore) responses.Response {
	// Read
	r.ParseForm()
	replyTo, err := strconv.ParseUint(r.FormValue("reply_to"), 10, 32)
	if err != nil {
		return responses.Status(400)
	}
	post := &models.Post{
		ReplyTo: uint32(replyTo),
		Body:    r.FormValue("body"),
	}

	// Validate
	post.Body = strings.TrimSpace(post.Body)
	if post.Body == "" || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}
	board, err := ds.GetThreadBoard(post.ReplyTo)
	if board == "" {
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
