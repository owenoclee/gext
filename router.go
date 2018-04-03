package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/owenoclee/gext-server/controllers"
)

func initRouter(db *sql.DB) *httprouter.Router {
	router := httprouter.New()
	router.PanicHandler = panicHandler

	router.POST("/posts", controllers.StorePost.Handler(db))
	router.OPTIONS("/posts", corsHandler)
	router.GET("/threads/:id", controllers.ShowThread.Handler(db))
	router.GET("/boards/:board/page/:page", controllers.ShowBoard.Handler(db))
	router.POST("/threads", controllers.StoreThread.Handler(db))
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
