package main

import _ "github.com/go-sql-driver/mysql"

func main() {
	db := initDB()
	defer db.Close()
	router := initRouter(db)
	serveHTTP(router)
}
