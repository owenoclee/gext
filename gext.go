package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/owenoclee/gext/datastore"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	env, err := initEnv()
	if err != nil {
		log.Fatal(err)
	}

	templates, err := initViews(env)
	if err != nil {
		log.Fatal(err)
	}

	ds, err := datastore.NewDatastore(env)
	if err != nil {
		log.Fatal(err)
	}
	defer ds.Close()

	router := initRouter(ds, templates, env)
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", env.Address(), env.Port()), router))
}
