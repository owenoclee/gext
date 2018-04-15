package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
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
	body := strings.TrimSpace(post.GetBody())
	if body == "" || len([]rune(body)) > 4000 {
		return responses.Status(422)
	}
	// Check the thread exists
	replyTo := post.GetReplyTo()
	thread := db.QueryRow("SELECT board FROM posts WHERE id = ? AND reply_to IS NULL", replyTo)
	var board sql.NullString
	if err := thread.Scan(&board); err == sql.ErrNoRows {
		return responses.Status(422)
	} else if err != nil {
		return responses.LogError(err)
	}

	// Store the post
	result, err := db.Exec(
		"INSERT INTO posts (reply_to, body) VALUES (?, ?)",
		replyTo,
		body,
	)
	if err != nil {
		return responses.LogError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return responses.LogError(err)
	}

	return responses.Created(fmt.Sprintf("/%v/thread/%v#%v", board.String, replyTo, id))
}
