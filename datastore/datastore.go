package datastore

import (
	"fmt"

	"github.com/owenoclee/gext/config"
	"github.com/owenoclee/gext/models"
)

type Datastore interface {
	StoreThread(thread models.Post) (uint32, error)
	GetThread(id uint32) (models.Thread, error)
	GetThreadBoard(id uint32) (string, error)
	StorePost(post models.Post) (uint32, error)
	GetPage(board string, pageNum uint32) (models.Page, error)
	Close() error
}

type datastoreFactory func(config.Env) (Datastore, error)

var registeredFactories map[string]datastoreFactory = map[string]datastoreFactory{
	"mysql": newMySQLDatastore,
}

func NewDatastore(env config.Env) (Datastore, error) {
	factory := registeredFactories[env.Read("DATASTORE")]
	if factory == nil {
		return nil, fmt.Errorf("Invalid DATASTORE: '%v'", env.Read("DATASTORE"))
	}
	return factory(env)
}
