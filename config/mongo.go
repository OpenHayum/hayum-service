package config

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

type Mongo struct {
	*mgo.Database
}

func NewMongoSession(url string, db string) (*Mongo, error) {
	if url == "" {
		return nil, fmt.Errorf("%s", "Mongo URL cannot be empty")
	}

	var session *mgo.Session

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	database := session.DB(db)

	return &Mongo{
		Database: database,
	}, nil
}
