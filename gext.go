package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/owenoclee/gext/config"
	"github.com/owenoclee/gext/datastore"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	var environment string
	for _, e := range os.Environ() {
		environment = environment + "\n" + e
	}
	envMap, err := godotenv.Unmarshal(environment)
	if err != nil {
		log.Fatal(err)
	}
	env := config.NewEnv(envMap)

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
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%v:%v", env.Read("ADDRESS"), env.Read("PORT")), router))
}
