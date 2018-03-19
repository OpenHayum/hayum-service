package models

import "gopkg.in/mgo.v2/bson"

// Artist holds the model of an Artist
type Artist struct {
	ID     bson.ObjectId `json:"_id" bson:"_id"`
	UserID string        `json:"userId" bson:"userId"`
	Name   string        `json:"name" bson:"name"`
}
