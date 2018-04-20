package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/owenoclee/gext-server/datastore"
)

func main() {
	db := initDB()
	defer db.Close()
	datastore.Initialize(db)
	router := initRouter(db)
	serveHTTP(router)
}
