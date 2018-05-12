package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/models"
	"github.com/owenoclee/gext/responses"
)

var StorePost Action = func(r *http.Request, _ httprouter.Params, ds datastore.Datastore) responses.Response {
	// Read
	postBinary, err := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	if err2 := proto.Unmarshal(postBinary, post); err != nil || err2 != nil {
		return responses.Status(400)
	}

	// Validate
	post.Body = strings.TrimSpace(post.GetBody())
	if post.Body == "" || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}
	board, err := ds.GetThreadBoard(post.GetReplyTo())
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
	return responses.Created(fmt.Sprintf("/%v/thread/%v#%v", board, post.GetReplyTo(), id))
}
