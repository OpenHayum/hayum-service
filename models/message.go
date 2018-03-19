package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Message contains Message model
type Message struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	SenderID    string        `json:"senderId" bson:"senderId"`
	RecipientID string        `json:"recipientId" bson:"recipientId"`
	Message     string        `json:"message" bson:"message"`
	Date        time.Time     `json:"date" bson:"date"`
}
