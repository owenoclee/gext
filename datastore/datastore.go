package datastore

import (
	"fmt"

	"github.com/owenoclee/gext-server/models"
)

type Datastore interface {
	StoreThread(thread *models.Post) (uint32, error)
	GetThread(id int64) (models.Thread, error)
	GetThreadBoard(id uint32) (string, error)
	StorePost(post *models.Post) (uint32, error)
	GetPage(board string, pageNum int64) (models.Page, error)
	Close() error
}

type datastoreFactory func(map[string]string) (Datastore, error)

var registeredFactories map[string]datastoreFactory = map[string]datastoreFactory{
	"mysql": newMySQLDatastoreFactory,
}

func NewDatastore(env map[string]string) (Datastore, error) {
	factory := registeredFactories[env["DATASTORE"]]
	if factory == nil {
		return nil, fmt.Errorf("Invalid DATASTORE: '%v'", env["DATASTORE"])
	}
	return factory(env)
}
