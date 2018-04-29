package models

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// User contains User model
type User struct {
	ID                 bson.ObjectId `json:"_id" bson:"_id"`
	ArtistID           string        `json:"artistId" bson:"artistId"`
	FullName           string        `json:"fullName" bson:"fullName"`
	UserName           string        `json:"userName" bson:"userName"`
	Email              string        `json:"email" bson:"email"`
	MobileNumber       string        `json:"mobileNumber" bson:"mobileNumber"`
	Password           string        `json:"password" bson:"password"`
	Role               string        `json:"role" bson:"role"`
	Verified           bool          `json:"verified" bson:"verified"`
	VerifiedAsAnArtist bool          `json:"verifiedAsAnArtist" bson:"verifiedAsAnArtist"`
	Otp                int32         `json:"otp" bson:"otp"`
	IsDeleted          bool          `json:"isDeleted" bson:"isDeleted"`
	OtpExpirationDate  time.Time     `json:"otpExpirationDate" bson:"otpExpirationDate"`
	CreatedDate        time.Time     `json:"createdDate" bson:"createdDate"`
	UpdatedDate        time.Time     `json:"updatedDate" bson:"updatedDate"`
	DeletedDate        time.Time     `json:"deletedDate" bson:"deletedDate"`
}
