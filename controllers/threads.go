package controllers

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/datastore"
	"github.com/owenoclee/gext-server/models"
	"github.com/owenoclee/gext-server/responses"
)

var boardRegex = regexp.MustCompile("^[a-z]{1,16}$")

var StoreThread Action = func(r *http.Request, _ httprouter.Params, ds datastore.Datastore) responses.Response {
	// Read
	postBinary, err := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	if err2 := proto.Unmarshal(postBinary, post); err != nil || err2 != nil {
		return responses.Status(400)
	}

	// Validate
	post.Board = strings.ToLower(strings.TrimSpace(post.GetBoard()))
	if !boardRegex.MatchString(post.Board) {
		return responses.Status(422)
	}
	post.Subject = strings.TrimSpace(post.GetSubject())
	post.Body = strings.TrimSpace(post.GetBody())
	if (post.Subject == "" && post.Body == "") || len([]rune(post.Subject)) > 32 || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}

	// Store
	id, err := ds.StoreThread(post)
	if err != nil {
		return responses.LogError(err)
	}
	return responses.Created(fmt.Sprintf("/%v/thread/%v", post.Board, id))
}

var ShowThread Action = func(_ *http.Request, p httprouter.Params, ds datastore.Datastore) responses.Response {
	// Read
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)

	// Validate
	if err != nil {
		return responses.Status(422)
	}

	// Show
	thread, err := ds.GetThread(id)
	if err != nil {
		return responses.LogError(err)
	} else if thread.GetPosts() == nil {
		return responses.Status(404)
	}
	return responses.Protobuf(&thread, 200)
}
