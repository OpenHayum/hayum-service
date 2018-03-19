package models

// Ads stores an ads info
type Ads struct {
	Type         string `json:"type" bson:"type"`
	Title        string `json:"title" bson:"title"`
	S3DocumentID string `json:"s3DocumentId" bson:"s3DocumentId"`
	Duration     string `json:"duration" bson:"duration"`
	PublisherID  string `json:"publisherId" bson:"publisherId"`
	Views        int64  `json:"views" bson:"views"`
	Clicks       int64  `json:"clicks" bson:"clicks"`
	Listens      int64  `json:"listens" bson:"listens"`
}
