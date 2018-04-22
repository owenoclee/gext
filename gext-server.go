package main

import (
	"log"
	"net/http"

	"github.com/owenoclee/gext-server/datastore"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	config := map[string]string{
		"DATASTORE":           "mysql",
		"DATASTORE_MYSQL_DSN": "root@/gext",
		"ADDRESS":             ":8080",
	}

	ds, err := datastore.NewDatastore(config)
	defer ds.Close()
	if err != nil {
		log.Fatal(err)
	}
	router := initRouter(ds)
	log.Fatal(http.ListenAndServe(config["ADDRESS"], router))
}
