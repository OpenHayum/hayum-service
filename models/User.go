package models

import (
	"gopkg.in/guregu/null.v3"
)

type User struct {
	Id           int64
	Email        string
	FirstName    string
	LastName     string
	Mobile       string
	Password     string
	Otp          int
	IsDeleted    int
	IsVerified   int
	OtpExpiresAt null.Time
	CreatedDate  null.Time
	DeletedDate  null.Time
	ModifiedDate null.Time
}
