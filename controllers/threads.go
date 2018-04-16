package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/drivers"
	"github.com/owenoclee/gext-server/models"
	"github.com/owenoclee/gext-server/responses"
)

var boardRegex = regexp.MustCompile("^[a-z]{1,16}$")

var StoreThread Action = func(r *http.Request, _ httprouter.Params, db *sql.DB) responses.Response {
	// Read the request
	postBinary, err := ioutil.ReadAll(r.Body)
	post := &models.Post{}
	if err2 := proto.Unmarshal(postBinary, post); err != nil || err2 != nil {
		return responses.Status(400)
	}

	// Validate the request
	post.Board = strings.ToLower(strings.TrimSpace(post.GetBoard()))
	if !boardRegex.MatchString(post.Board) {
		return responses.Status(422)
	}
	post.Subject = strings.TrimSpace(post.GetSubject())
	post.Body = strings.TrimSpace(post.GetBody())
	if (post.Subject == "" && post.Body == "") || len([]rune(post.Subject)) > 32 || len([]rune(post.Body)) > 4000 {
		return responses.Status(422)
	}

	id, err := drivers.StoreThread(post)
	if err != nil {
		return responses.LogError(err)
	}
	return responses.Created(fmt.Sprintf("/%v/thread/%v", post.Board, id))
}

var ShowThread Action = func(r *http.Request, p httprouter.Params, db *sql.DB) responses.Response {
	// Validate parameters
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		return responses.Status(422)
	}

	thread, err := drivers.GetThread(id)
	if err != nil {
		return responses.LogError(err)
	} else if thread.GetPosts() == nil {
		return responses.Status(404)
	}
	return responses.Protobuf(&thread, 200)
}
