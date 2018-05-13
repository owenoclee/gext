package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/config"
	"github.com/owenoclee/gext/controllers"
	"github.com/owenoclee/gext/datastore"
)

func initRouter(ds datastore.Datastore, t *template.Template, env config.Env) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = panicHandler

	router.POST("/posts", controllers.StorePost.Handler(ds, t))
	router.OPTIONS("/posts", corsHandler)
	router.GET("/start-thread", controllers.CreateThread.Handler(ds, t))
	router.GET("/threads/:id", controllers.ShowThread.Handler(ds, t))
	router.GET("/boards/:board/page/:page", controllers.ShowBoard.Handler(ds, t))
	router.POST("/threads", controllers.StoreThread.Handler(ds, t))
	router.OPTIONS("/threads", corsHandler)
	router.ServeFiles("/static/*filepath", http.Dir(env.Read("PUBLIC_PATH")))

	return router
}

func corsHandler(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "POST")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
}

func panicHandler(w http.ResponseWriter, r *http.Request, err interface{}) {
	log.Printf("panic handling http %v request to '%v':\n%v\n", r.Method, r.RequestURI, err)
	w.WriteHeader(500)
}
