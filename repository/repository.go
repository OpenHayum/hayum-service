package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
)

type Repository struct {
	*config.Mongo
}

func NewRepository(mongo *Mongo) *Repository {
	return &Repository{mongo}
}
