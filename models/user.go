package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// RoleAdmin can do everything
	RoleAdmin = "ADMIN"

	// RoleModerator can act as a moderator for uploaded contents
	RoleModerator = "MODERATOR"

	// RoleUser acts as a normal user
	RoleUser = "USER"

	// RoleArtist will have privileges of an artist
	RoleArtist = "ARTIST"
)

// User contains User model
type User struct {
	ID                        bson.ObjectId `json:"_id" bson:"_id"`
	ArtistID                  string        `json:"artistId" bson:"artistId"`
	FullName                  string        `json:"fullName" bson:"fullName"`
	Email                     string        `json:"email" bson:"email"`
	MobileNumber              string        `json:"mobileNumber" bson:"mobileNumber"`
	Otp                       rune          `json:"otp" bson:"otp"`
	Salt                      string        `json:"salt" bson:"salt"`
	Role                      string        `json:"role" bson:"role"`
	IsUserVerified            bool          `json:"isUserVerified" bson:"isUserVerified"`
	HasUserVerifiedAsAnArtist bool          `json:"hasUserVerifiedAsAnArtist" bson:"hasUserVerifiedAsAnArtist"`
	Meta                      userMeta      `json:"meta" bson:"meta"`
	CreatedDate               time.Time     `json:"createdDate" bson:"createdDate"`
	UpdatedDate               time.Time     `json:"updatedDate" bson:"updatedDate"`
	DeletedDate               time.Time     `json:"deletedDate" bson:"deletedDate"`
}

type userMeta struct {
	Size                  string `json:"size" bson:"size"`
	Downloads             int    `json:"downloads" bson:"downloads"`
	Views                 int    `json:"views" bson:"views"`
	NumberOfFavorites     int    `json:"numberOfFavorites" bson:"numberOfFavorites"`
	NumberOfItemsUploaded int    `json:"numberOfItemsUploaded" bson:"numberOfItemsUploaded"`
	ProfileImageLink      string `json:"profileImageLink" bson:"profileImageLink"`
	CoverImageLink        string `json:"coverImageLink" bson:"coverImageLink"`
}
