package models

import (
	"gopkg.in/mgo.v2/bson"
)

// Account stores all the user metas
type Account struct {
	ID                     bson.ObjectId
	UserID                 string `json:"userId" bson:"userId"`
	Downloads              int    `json:"downloads" bson:"downloads"`
	Views                  int    `json:"views" bson:"views"`
	AllowedDownloadsPerDay int    `json:"allowedDownloadsPerDay" bson:"allowedDownloadsPerDay"`
	CurrentNumOfDownloads  int    `json:"currentNumOfDownloads" bson:"currentNumOfDownloads"`
	TotalNumOfDownloads    int64  `json:"totalNumOfDownloads" bson:"totalNumOfDownloads"`
	IsPremiumUser          bool   `json:"isPremiumUser" bson:"isPremiumUser"`
	NumberOfFavorites      int    `json:"numberOfFavorites" bson:"numberOfFavorites"`
	NumberOfItemsUploaded  int    `json:"numberOfItemsUploaded" bson:"numberOfItemsUploaded"`
	ProfileImageLink       string `json:"profileImageLink" bson:"profileImageLink"`
	CoverImageLink         string `json:"coverImageLink" bson:"coverImageLink"`
}
