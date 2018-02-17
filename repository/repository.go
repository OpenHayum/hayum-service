package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"gopkg.in/mgo.v2"
)

// Repository maintains a mgo.Collection
type Repository struct {
	*mgo.Collection
}

// NewRepository creates a new Repository
func NewRepository(mongo *config.Mongo, collectionName string) *Repository {
	return &Repository{mongo.C(collectionName)}
}
