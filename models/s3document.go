package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// S3Document holds the structure of a S3Document
type S3Document struct {
	ID               bson.ObjectId `json:"_id" bson:"_id"`
	OriginalFileName string        `json:"originalFileName" bson:"originalFileName"`
	Key              string        `json:"key" bson:"key"` // originalFileName + hash
	Bucket           string        `json:"bucket" bson:"bucket"`
	IsDeleted        bool          `json:"isDeleted" bson:"isDeleted"`
	CreatedDate      time.Time     `json:"createdDate" bson:"createdDate"`
	ModifiedDate     time.Time     `json:"modifiedDate" bson:"modifiedDate"`
	DeletedDate      time.Time     `json:"deletedDate" bson:"deletedDate"`
}
