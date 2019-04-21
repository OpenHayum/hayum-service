package models

import (
	"gopkg.in/guregu/null.v3"
	"time"
)

type User struct {
	Id           int
	Email        string
	FirstName    string
	LastName     string
	Mobile       string
	Password     string
	IsDeleted    int
	IsVerified   int
	OtpExpiresAt null.Time
	CreatedDate  time.Time
	DeletedDate  null.Time
	ModifiedDate null.Time
}
