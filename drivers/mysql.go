package drivers

import (
	"database/sql"
	"time"

	"github.com/owenoclee/gext-server/models"
)

var db *sql.DB

func Initialize(_db *sql.DB) {
	db = _db
}

func GetPage(board string, pageNum int64) (models.Page, error) {
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
		board,
		(pageNum-1)*15,
		15,
	)
	defer posts.Close()
	if err == sql.ErrNoRows {
		return models.Page{}, nil
	} else if err != nil {
		return models.Page{}, err
	}

	// Decode database response into a thread (collection of posts)
	page := models.Page{}
	curThread := &models.Thread{}
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
			return models.Page{}, err
		}
		postResponse := models.Post{
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
				page.Threads = append(page.Threads, curThread)
			}
			curThread = &models.Thread{}
		}
		curThread.Posts = append(curThread.Posts, &postResponse)
	}
	if len(curThread.Posts) > 0 {
		page.Threads = append(page.Threads, curThread)
	}

	return page, nil
}

func GetThreadBoard(id uint32) (string, error) {
	thread := db.QueryRow("SELECT board FROM posts WHERE id = ? AND reply_to IS NULL", id)
	var board sql.NullString
	if err := thread.Scan(&board); err == sql.ErrNoRows {
		return "", nil
	} else if err != nil {
		return "", err
	}
	return board.String, nil
}

func StorePost(post *models.Post) (uint32, error) {
	// Store the post
	result, err := db.Exec(
		"INSERT INTO posts (reply_to, body) VALUES (?, ?)",
		post.GetReplyTo(),
		post.GetBody(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint32(id), err
}

func StoreThread(thread *models.Post) (uint32, error) {
	// Store the post
	result, err := db.Exec(
		"INSERT INTO posts (board, subject, body) VALUES (?, ?, ?)",
		thread.GetBoard(),
		thread.GetSubject(),
		thread.GetBody(),
	)
	if err != nil {
		return 0, err
	}
	id, err := result.LastInsertId()
	return uint32(id), err
}

func GetThread(id int64) (models.Thread, error) {
	// Fetch all posts in thread, incl the thread itself
	posts, err := db.Query(
		"SELECT id, reply_to, board, subject, body, created_at FROM posts WHERE id = ? OR reply_to = ?",
		id,
		id,
	)
	defer posts.Close()
	if err == sql.ErrNoRows {
		return models.Thread{}, nil
	} else if err != nil {
		return models.Thread{}, err
	}

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
			return models.Thread{}, err
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

	return *thread, nil
}
