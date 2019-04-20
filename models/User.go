package models

import "time"

type User struct {
	CreatedDate  time.Time
	DeletedDate  time.Time
	Email        string
	FirstName    string
	ID           int
	IsDeleted    int
	IsVerified   int
	LastName     string
	Mobile       string
	ModifiedDate time.Time
	OtpExpiresAt time.Time
	Password     string
}
