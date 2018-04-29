package db

import (
	"fmt"

	"gopkg.in/mgo.v2"
)

// Session exports mgo.Session
var MongoSession *mgo.Session

// NewMongoSession creates a mongo session
func NewMongoSession(url string, db string) error {
	if url == "" {
		return fmt.Errorf("%s", "Mongo URL cannot be empty")
	}

	session, err := mgo.Dial(url)
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	MongoSession = session

	return nil
}
