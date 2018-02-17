package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"gopkg.in/mgo.v2"
)

type Repository struct {
	*mgo.Collection
}

func NewRepository(mongo *config.Mongo, collectionName string) *Repository {
	return &Repository{mongo.C(collectionName)}
}
