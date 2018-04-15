package controllers

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/julienschmidt/httprouter"
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
	board := strings.ToLower(strings.TrimSpace(post.GetBoard()))
	if !boardRegex.MatchString(board) {
		return responses.Status(422)
	}
	subject, body := strings.TrimSpace(post.GetSubject()), strings.TrimSpace(post.GetBody())
	if (subject == "" && body == "") || len([]rune(subject)) > 32 || len([]rune(body)) > 4000 {
		return responses.Status(422)
	}

	// Store the post
	result, err := db.Exec(
		"INSERT INTO posts (board, subject, body) VALUES (?, ?, ?)",
		board,
		subject,
		body,
	)
	if err != nil {
		return responses.LogError(err)
	}
	id, err := result.LastInsertId()
	if err != nil {
		return responses.LogError(err)
	}

	return responses.Created(fmt.Sprintf("/%v/thread/%v", board, id))
}

var ShowThread Action = func(r *http.Request, p httprouter.Params, db *sql.DB) responses.Response {
	// Validate parameters
	id, err := strconv.ParseInt(p.ByName("id"), 10, 64)
	if err != nil {
		return responses.Status(422)
	}

	// Fetch all posts in thread, incl the thread itself
	posts, err := db.Query(
		"SELECT id, reply_to, board, subject, body, created_at FROM posts WHERE id = ? OR reply_to = ?",
		id,
		id,
	)
	if err == sql.ErrNoRows {
		return responses.Status(404)
	} else if err != nil {
		return responses.LogError(err)
	}
	defer posts.Close()

	// Decode database response into a thread (collection of posts)
	thread := &models.Thread{}
	for posts.Next() {
		var (
			pID       int64
			replyTo   sql.NullInt64
			board     sql.NullString
			subject   sql.NullString
			body      sql.NullString
			createdAt time.Time
		)
		if err := posts.Scan(&pID, &replyTo, &board, &subject, &body, &createdAt); err != nil {
			return responses.LogError(err)
		}
		post := models.Post{
			Id:        uint32(pID),
			ReplyTo:   uint32(replyTo.Int64),
			Board:     board.String,
			Subject:   subject.String,
			Body:      body.String,
			CreatedAt: uint32(createdAt.Unix()),
		}
		thread.Posts = append(thread.Posts, &post)
	}

	return responses.Protobuf(thread, 200)
}
