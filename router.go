package main

import (
	"html/template"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/owenoclee/gext/config"
	"github.com/owenoclee/gext/controllers"
	"github.com/owenoclee/gext/datastore"
	"github.com/owenoclee/gext/responses"
	"goji.io"
	"goji.io/pat"
)

func initRouter(ds datastore.Datastore, t *template.Template, env config.Env) *goji.Mux {
	mux := goji.NewMux()

	mux.Use(panicHandler)
	mux.Handle(pat.Get("/"), http.RedirectHandler("/general", 302))
	mux.Handle(pat.Get("/static/*"), http.StripPrefix("/static", http.FileServer(http.Dir(env.PublicPath()))))
	mux.Handle(pat.Get("/create-thread"), controllers.CreateThread.Handler(ds, t))
	mux.Handle(pat.Get("/:board"), controllers.ShowBoard.Handler(ds, t))
	mux.Handle(pat.Get("/:board/:page"), controllers.ShowBoard.Handler(ds, t))
	mux.Handle(pat.Get("/:board/thread/:id"), controllers.ShowThread.Handler(ds, t))
	mux.Handle(pat.Post("/:board/thread/:id/post"), controllers.StorePost.Handler(ds, t))
	mux.Handle(pat.Post("/threads"), controllers.StoreThread.Handler(ds, t))

	return mux
}

func panicHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("error: %v\n\nrequest: %v\n\nstack trace: %v\n", err, r, string(debug.Stack()))
				responses.Status(500).Write(w)
			}
		}()

		h.ServeHTTP(w, r)
	})
}
