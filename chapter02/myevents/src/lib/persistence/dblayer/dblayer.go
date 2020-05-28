package dblayer

import (
	/*
		"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence"
		"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence/mongolayer"
	*/
	"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence"
	"github.com/Cloud-Native-programming-with-Golang/chapter02/myevents/src/lib/persistence/mongolayer"
)

type DBTYPE string

const (
	MONGODB  DBTYPE = "mongodb"
	DYBAMODB DBTYPE = "dynamodb"
)

func NewPersistenceLayer(options DBTYPE, connection string) (persistence.DatabaseHandler, error) {
	switch options {
	case MONGODB:
		return mongolayer.NewMongoDBLayer(connection)
	}
	return nil, nil
}
