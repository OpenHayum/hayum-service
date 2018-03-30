package repository

import (
	"bitbucket.org/hayum/hayum-service/db"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// BaseRepository defines all default repository methods
type BaseRepository interface {
	Save(model interface{}) error
	GetByID(id string) (interface{}, error)
}

// MongoRepository maintains a mgo.Collection which persist data in the database
type MongoRepository struct {
	collection *mgo.Collection
}

// NewRepository creates a new Repository
func NewRepository(mongo *db.Mongo, collectionName string) *MongoRepository {
	session := mongo.Session.Copy()
	collection := session.DB(mongo.Db).C(collectionName)
	return &MongoRepository{collection}
}

// Save implements BaseRepository Save
func (mr *MongoRepository) Save(model interface{}) error {
	err := mr.collection.Insert(model)
	return err
}

// GetByID implements BaseRepository GetByID
func (mr *MongoRepository) GetByID(id string, model interface{}) error {
	return mr.collection.FindId(bson.ObjectIdHex(id)).One(model)
}
