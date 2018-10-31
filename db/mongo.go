package db

import (
	"fmt"
	"log"

	"gopkg.in/mgo.v2"
)

// Session exports mgo.Session
var MongoSession *mgo.Session

// NewMongoSession creates a mongo session
func NewMongoSession(url string, db string) error {
	log.Printf("Connecting, URL: %s", url)
	if url == "" {
		return fmt.Errorf("%s", "Mongo URL cannot be empty")
	}

	session, err := mgo.Dial(url)
	if err != nil {
		fmt.Print("Unable to connect with mongo URL:", url)
		return fmt.Errorf("%s", err)
	}

	MongoSession = session

	return nil
}
