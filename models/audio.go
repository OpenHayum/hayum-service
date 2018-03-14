package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	HyCategorySong       = "SONG"
	HyCategorySumangLila = "SUMANG_LILA"
	HyCategoryRadioLila  = "RADIO_LILA"
)

// Audio contains Audio model
type Audio struct {
	ID           bson.ObjectId `json:"_id" bson:"_id"`
	Name         string        `json:"name" bson:"name"`
	ArtistID     string        `json:"artistId" bson:"artistId"`
	IsOldSong    bool          `json:"isOldSong" bson:"isOldSong"`
	Album        string        `json:"album" bson:"album"`
	Category     string        `json:"category" bson:"category"`
	Thumbnail    string        `json:"thumbnail" bson:"thumbnail"`
	Status       string        `json:"status" bson:"status"`
	S3DocumentID string        `json:"s3DocumentID" bson:"s3DocumentID"`
	Meta         audioMeta     `json:"meta" bson:"meta"`
	UploadedBy   string        `json:"uploadedBy" bson:"uploadedBy"`
	ModeratedBy  string        `json:"moderatedBy" bson:"moderatedBy"`
	CreatedDate  time.Time     `json:"createdDate" bson:"createdDate"`
	UpdatedDate  time.Time     `json:"updatedDate" bson:"updatedDate"`
	DeletedDate  time.Time     `json:"deletedDate" bson:"deletedDate"`
}

type audioMeta struct {
	Size              rune  `json:"size" bson:"size"`
	Downloads         int64 `json:"downloads" bson:"downloads"`
	Views             int64 `json:"views" bson:"views"`
	NumberOfFavorites int64 `json:"numberOfFavorites" bson:"numberOfFavorites"`
}
