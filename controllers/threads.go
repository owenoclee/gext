package controllers

import (
	"fmt"
	"html/template"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/models"
	"github.com/owenoclee/gext/responses"
	"goji.io/pat"
)

var boardRegex = regexp.MustCompile("^[a-z]{1,16}$")

var CreateThread Action = func(_ *http.Request, _ datastore.Datastore, t *template.Template) responses.Response {
	return responses.View(t.Lookup("create-thread.html"), responses.ViewData{Title: "create thread - gext"})
}

var StoreThread Action = func(r *http.Request, ds datastore.Datastore, t *template.Template) responses.Response {
	// Read
	r.ParseForm()
	post := models.Post{
		Board:   r.FormValue("board"),
		Subject: r.FormValue("subject"),
		Body:    r.FormValue("body"),
	}

	// Validate
	post.Board = strings.ToLower(strings.TrimSpace(post.Board))
	if !boardRegex.MatchString(post.Board) {
		return responses.Status(422)
	}
	post.Subject = strings.TrimSpace(post.Subject)
	post.Body = strings.TrimSpace(post.Body)
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

var ShowThread Action = func(r *http.Request, ds datastore.Datastore, t *template.Template) responses.Response {
	// Read
	id64, err := strconv.ParseUint(pat.Param(r, "id"), 10, 32)

	// Validate
	if err != nil {
		return responses.Status(422)
	}
	id := uint32(id64)

	// Show
	thread, err := ds.GetThread(id)
	if err != nil {
		return responses.LogError(err)
	} else if len(thread.Posts) == 0 {
		return responses.Status(404)
	}
	return responses.View(t.Lookup("thread.html"), responses.ViewData{
		Title: fmt.Sprintf("/%v/ thread - gext", thread.Board()),
		Data:  thread,
	})
}
