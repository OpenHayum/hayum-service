package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type S3Document struct {
	ID               bson.ObjectId `json:"_id" bson:"_id"`
	OriginalFileName string        `json:"originalFileName" bson:"originalFileName"`
	IsDeleted        bool          `json:"isDeleted" bson:"isDeleted"`
	CreatedDate      time.Time     `json:"createdDate" bson:"createdDate"`
	DeletedDate      time.Time     `json:"deletedDate" bson:"deletedDate"`
}
