package models

import "time"

type Account struct {
	AccountImageLink string
	CoverImageLink   string
	CreatedDate      time.Time
	DeletedDate      time.Time
	Id               int64
	IsDeleted        int
	IsPremium        int
	ModifiedDate     time.Time
	UserId           int64
	ViewsNum         int
}
