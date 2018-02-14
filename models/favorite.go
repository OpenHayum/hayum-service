package models

import "gopkg.in/mgo.v2/bson"

type Favorite struct {
	Id       bson.ObjectId `json:"_id" bson:"_id"`
	UserId   string        `json:"userId" bson:"userId"`
	ItemId   string        `json:"itemId" bson:"itemId"`
	ItemName string        `json:"itemName" bson:"itemName"`
	Artist   string        `json:"artist" bson:"artist"`
}
