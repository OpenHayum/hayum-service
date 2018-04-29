package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"gopkg.in/mgo.v2/bson"
)

type S3DocumentRepositorer interface {
	MongoRepositorer
	CreateNewS3Document(doc *models.S3Document) error
}

// S3DocumentRepository embeds a Repository
type S3DocumentRepository struct {
	*MongoRepository
}

// NewS3DocumentRepository creates a new S3DocumentRepository
func NewS3DocumentRepository() *S3DocumentRepository {
	return &S3DocumentRepository{NewMongoRepository(config.CollectionS3Document)}
}

// CreateNewS3Document creates a new S3 Document
func (r *S3DocumentRepository) CreateNewS3Document(doc *models.S3Document) error {
	doc.ID = bson.NewObjectId()
	err := r.Save(&doc)
	return err
}
