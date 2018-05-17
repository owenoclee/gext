package main

import (
	"html/template"
	"log"
	"net/http"

	"goji.io"
	"goji.io/pat"

	"github.com/owenoclee/gext/config"
	"github.com/owenoclee/gext/controllers"
	"github.com/owenoclee/gext/datastore"
)

func initRouter(ds datastore.Datastore, t *template.Template, env config.Env) *goji.Mux {
	mux := goji.NewMux()

	mux.Handle(pat.Get("/static/*"), http.StripPrefix("/static", http.FileServer(http.Dir(env.PublicPath()))))
	mux.Handle(pat.Post("/posts"), controllers.StorePost.Handler(ds, t))
	mux.Handle(pat.Get("/start-thread"), controllers.CreateThread.Handler(ds, t))
	mux.Handle(pat.Get("/threads/:id"), controllers.ShowThread.Handler(ds, t))
	mux.Handle(pat.Get("/boards/:board"), controllers.ShowBoard.Handler(ds, t))
	mux.Handle(pat.Get("/boards/:board/page/:page"), controllers.ShowBoard.Handler(ds, t))
	mux.Handle(pat.Post("/threads"), controllers.StoreThread.Handler(ds, t))

	return mux
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("panic handling http %v request to '%v':\n%v\n", r.Method, r.RequestURI, err)
	w.WriteHeader(500)
}
