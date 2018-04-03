package controllers

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"github.com/owenoclee/gext-server/network"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/responses"
)

var ShowBoard Action = func(r *http.Request, p httprouter.Params, db *sql.DB) responses.Response {
	// Validate parameters
	page, err := strconv.ParseInt(p.ByName("page"), 10, 64)
	if err != nil {
		return responses.Status(422)
	}
	if page < 1 {
		page = 1
	}

	// Fetch latest threads + posts
	posts, err := db.Query(
		`SELECT posts.* FROM posts
			JOIN (
				SELECT * FROM latest_threads
				WHERE board = ?
				ORDER BY bumped_at DESC
				LIMIT ?, ?
			) AS latest_threads ON latest_threads.thread_id = posts.id
			OR latest_threads.thread_id = posts.reply_to
			ORDER BY bumped_at DESC, created_at ASC`,
		p.ByName("board"),
		(page-1)*15,
		15,
	)
	if err == sql.ErrNoRows {
		return responses.Status(404)
	} else if err != nil {
		return responses.LogError(err)
	}
	defer posts.Close()

	// Decode database response into a thread (collection of posts)
	pageResponse := network.PageResponse{}
	curThread := &network.ThreadResponse{}
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
		postResponse := network.PostResponse{
			Id:        uint32(pID),
			ReplyTo:   uint32(replyTo.Int64),
			Board:     board.String,
			Subject:   subject.String,
			Body:      body.String,
			CreatedAt: uint32(createdAt.Unix()),
		}

		// if this post is a thread
		if postResponse.GetReplyTo() == 0 {
			if len(curThread.Posts) > 0 {
				pageResponse.Threads = append(pageResponse.Threads, curThread)
			}
			curThread = &network.ThreadResponse{}
		}
		curThread.Posts = append(curThread.Posts, &postResponse)
	}
	if len(curThread.Posts) > 0 {
		pageResponse.Threads = append(pageResponse.Threads, curThread)
	}

	return responses.Protobuf(&pageResponse, 200)
}
