package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Mongo holds mongo session and Db
type Mongo struct {
	*mgo.Session
	Db string
}

// NewMongoSession creates a mongo session
func NewMongoSession(url string, db string) (*Mongo, error) {
	if url == "" {
		return nil, fmt.Errorf("%s", "Mongo URL cannot be empty")
	}

	session, err := mgo.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("%s", err)
	}

	return &Mongo{
		session,
		db,
	}, nil
}
