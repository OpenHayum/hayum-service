package models

import "gopkg.in/mgo.v2/bson"

type Favorite struct {
	ID       bson.ObjectId `json:"_id" bson:"_id"`
	UserID   string        `json:"userId" bson:"userId"`
	ItemID   string        `json:"itemId" bson:"itemId"`
	ItemName string        `json:"itemName" bson:"itemName"`
	Artist   string        `json:"artist" bson:"artist"`
}
