package config

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"

	"github.com/go-sql-driver/mysql"
)

type Env interface {
	SetAddress(string) error
	SetPort(string) error
	SetDatastore(string) error
	SetDatastoreMySqlDsn(string) error
	SetViewsPath(string) error
	SetPublicPath(string) error
	Address() string
	Port() string
	Datastore() string
	DatastoreMySqlDsn() string
	ViewsPath() string
	PublicPath() string
}

type env struct {
	address           string
	port              string
	datastore         string
	datastoreMySqlDsn string
	viewsPath         string
	publicPath        string
}

func NewEnv() Env {
	return &env{}
}

func (e *env) SetAddress(address string) error {
	if isEmpty(address) || net.ParseIP(address) == nil {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_ADDRESS'", address)
	}
	e.address = address
	return nil
}

func (e *env) SetPort(port string) error {
	if _, err := strconv.ParseUint(port, 10, 16); isEmpty(port) || err != nil {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_PORT'", port)
	}
	e.port = port
	return nil
}

func (e *env) SetDatastore(datastore string) error {
	if isEmpty(datastore) || datastore != "mysql" {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_DATASTORE'", datastore)
	}
	e.datastore = datastore
	return nil
}

func (e *env) SetDatastoreMySqlDsn(dsn string) error {
	if _, err := mysql.ParseDSN(dsn); isEmpty(dsn) || err != nil {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_DATASTORE_MYSQL_DSN'", dsn)
	}
	e.datastoreMySqlDsn = dsn
	return nil
}

func (e *env) SetViewsPath(path string) error {
	dir, err := os.Stat(path)
	if isEmpty(path) || err != nil {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_VIEWS_PATH'", path)
	} else if !dir.IsDir() {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_VIEWS_PATH' (not a directory)", path)
	}
	e.viewsPath = path
	return nil
}

func (e *env) SetPublicPath(path string) error {
	dir, err := os.Stat(path)
	if isEmpty(path) || err != nil {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_PUBLIC_PATH'", path)
	} else if !dir.IsDir() {
		return fmt.Errorf("config: '%v' is not a valid value for 'GEXT_PUBLIC_PATH' (not a directory)", path)
	}
	e.publicPath = path
	return nil
}

func (e *env) Address() string           { return e.address }
func (e *env) Port() string              { return e.port }
func (e *env) Datastore() string         { return e.datastore }
func (e *env) DatastoreMySqlDsn() string { return e.datastoreMySqlDsn }
func (e *env) ViewsPath() string         { return e.viewsPath }
func (e *env) PublicPath() string        { return e.publicPath }

func isEmpty(s string) bool {
	return strings.Trim(s, " ") == ""
}
