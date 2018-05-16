package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/owenoclee/gext/config"
)

type multiError []error

func newMultiError(es []error) error {
	for _, e := range es {
		if e != nil {
			return multiError(es)
		}
	}
	return nil
}

func (e multiError) Error() string {
	var errStr string
	for _, err := range e {
		if err != nil {
			errStr = errStr + "\n" + err.Error()
		}
	}
	return errStr
}

func initEnv() (config.Env, error) {
	var es []error
	m := getEnvMap()
	e := config.NewEnv()

	es = append(es, e.SetAddress(m["GEXT_ADDRESS"]))
	es = append(es, e.SetPort(m["GEXT_PORT"]))
	es = append(es, e.SetDatastore(m["GEXT_DATASTORE"]))
	es = append(es, e.SetDatastoreMySqlDsn(m["GEXT_DATASTORE_MYSQL_DSN"]))
	es = append(es, e.SetViewsPath(m["GEXT_VIEWS_PATH"]))
	es = append(es, e.SetPublicPath(m["GEXT_PUBLIC_PATH"]))

	return e, newMultiError(es)
}

func getEnvMap() map[string]string {
	var environment string
	for _, e := range os.Environ() {
		environment = environment + "\n" + e
	}
	envMap, err := godotenv.Unmarshal(environment)
	if err != nil {
		log.Fatal(err)
	}
	return envMap
}
