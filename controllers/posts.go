package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/datastore"
	"github.com/owenoclee/gext-server/models"
	"github.com/owenoclee/gext-server/responses"
)

var StorePost Action = func(r *http.Request, _ httprouter.Params, db *sql.DB) responses.Response {
	// Read the request
	postBinary, err := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	if err2 := proto.Unmarshal(postBinary, post); err != nil || err2 != nil {
		return responses.Status(400)
	}

	// Validate the request
	post.Body = strings.TrimSpace(post.GetBody())
	if post.Body == "" || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}
	// Check the thread exists
	board, err := datastore.GetThreadBoard(post.GetReplyTo())
	if board == "" {
		if err != nil {
			return responses.LogError(err)
		}
		return responses.Status(422)
	}

	id, err := datastore.StorePost(post)
	if err != nil {
		return responses.LogError(err)
	}
	return responses.Created(fmt.Sprintf("/%v/thread/%v#%v", board, post.GetReplyTo(), id))
}
