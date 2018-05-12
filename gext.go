package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/joho/godotenv"
	"github.com/owenoclee/gext/datastore"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env, err := godotenv.Read()
	if err != nil {
		log.Fatal(err)
	}

	ds, err := datastore.NewDatastore(env)
	defer ds.Close()
	if err != nil {
		log.Fatal(err)
	}
	router := initRouter(ds)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", env["ADDRESS"], env["PORT"]), router))
}
