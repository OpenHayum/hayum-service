package models

import "time"

type Account struct {
	AccountImageLink string
	CoverImageLink   string
	CreatedDate      time.Time
	DeletedDate      time.Time
	ID               int
	IsDeleted        int
	IsPremium        int
	ModifiedDate     time.Time
	UserID           int
	ViewsNum         int
}
