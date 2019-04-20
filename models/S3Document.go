package models

import "time"

type S3Document struct {
	Bucket           string
	CreatedDate      time.Time
	DeletedDate      time.Time
	ID               int
	IsDeleted        int
	ModifiedDate     time.Time
	OriginalFileName string
	S3Key            string
}
