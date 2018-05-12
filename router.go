package main

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext/controllers"
	"github.com/owenoclee/gext/datastore"
)

func initRouter(ds datastore.Datastore) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = panicHandler

	router.POST("/posts", controllers.StorePost.Handler(ds))
	router.OPTIONS("/posts", corsHandler)
	router.GET("/threads/:id", controllers.ShowThread.Handler(ds))
	router.GET("/boards/:board/page/:page", controllers.ShowBoard.Handler(ds))
	router.POST("/threads", controllers.StoreThread.Handler(ds))
	router.OPTIONS("/threads", corsHandler)

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
