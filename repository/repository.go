package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"gopkg.in/mgo.v2"
)

// BaseRepository defines all default repository methods
type BaseRepository interface {
	Save(model interface{}) error
	GetByID(id string) (interface{}, error)
}

// Repository maintains a mgo.Collection
type Repository struct {
	collection *mgo.Collection
}

// NewRepository creates a new Repository
func NewRepository(mongo *config.Mongo, collectionName string) *Repository {
	return &Repository{mongo.C(collectionName)}
}

// Save implements BaseRepository Save
func (r *Repository) Save(model interface{}) error {
	err := r.collection.Insert(model)
	return err
}

// GetByID implements BaseRepository GetByID
func (r *Repository) GetByID(id string) (interface{}, error) {
	var model interface{}
	err := r.collection.FindId(id).One(model)
	return &model, err
}
