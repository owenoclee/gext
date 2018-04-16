package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/owenoclee/gext-server/drivers"
)

func main() {
	db := initDB()
	defer db.Close()
	drivers.Initialize(db)
	router := initRouter(db)
	serveHTTP(router)
}
