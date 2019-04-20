package models

import "time"

type Account struct {
	AccountImageLink  string
	CoverImageLink    string
	CreatedDate       time.Time
	DeletedDate       time.Time
	FavoritesNum      int
	ID                int
	IsDeleted         int
	IsPremium         int
	ModifiedDate      time.Time
	TracksUploadedNum int
	UserID            int
	ViewsNum          int
}
