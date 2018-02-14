package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type User struct {
	ID          bson.ObjectId `json:"_id" bson:"_id"`
	ArtistID    string        `json:"artistId" bson:"artistId"`
	FullName    string        `json:"fullname" bson:"fullname"`
	IsOldSong   bool          `json:"isOldSong" bson:"isOldSong"`
	Album       string        `json:"album" bson:"album"`
	Category    string        `json:"category" bson:"category"`
	Thumbnail   string        `json:"thumbnail" bson:"thumbnail"`
	Status      string        `json:"status" bson:"status"`
	Document    S3Document    `json:"s3Document" bson:"s3Document"`
	Meta        userMeta      `json:"meta" bson:"meta"`
	UploadedBy  string        `json:"uploadedBy" bson:"uploadedBy"`
	ModeratedBy string        `json:"moderatedBy" bson:"moderatedBy"`
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
