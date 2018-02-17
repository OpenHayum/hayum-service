package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	ArtistID    string        `json:"artistId" bson:"artistId"`
	FullName    string        `json:"fullname" bson:"fullname"`
	Email       string        `json:"email" bson:"email"`
	Status      string        `json:"status" bson:"status"`
	Document    S3Document    `json:"s3Document" bson:"s3Document"`
	Meta        userMeta      `json:"meta" bson:"meta"`
	CreatedDate time.Time     `json:"createdDate" bson:"createdDate"`
	UpdatedDate time.Time     `json:"updatedDate" bson:"updatedDate"`
	DeletedDate time.Time     `json:"deletedDate" bson:"deletedDate"`
}

type userMeta struct {
	Size              rune  `json:"size" bson:"size"`
	Downloads         int64 `json:"downloads" bson:"downloads"`
	Views             int64 `json:"views" bson:"views"`
	NumberOfFavorites int64 `json:"numberOfFavorites" bson:"numberOfFavorites"`
}
