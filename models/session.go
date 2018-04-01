package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Session stores user sesssion
type Session struct {
	ID        bson.ObjectId
	UserID    string    `json:"userId" bson:"userId"`
	ExpiresAt time.Time `json:"expiresAt" bson:"expiresAt"`
}
