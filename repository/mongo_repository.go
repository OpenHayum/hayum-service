package repository

import (
	"bitbucket.org/hayum/hayum-service/db"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// MongoRepositorer implements Repositorer and some more mongo methods
type MongoRepositorer interface {
	Repositorer
	Count() (int, error)
}

// MongoRepository maintains a mgo.Collection which persist data in the database
type MongoRepository struct {
	collection *mgo.Collection
}

// NewMongoRepository creates a new Repository
func NewMongoRepository(mongo *db.Mongo, collectionName string) *MongoRepository {
	session := mongo.Session.Copy()
	collection := session.DB(mongo.Db).C(collectionName)
	return &MongoRepository{collection}
}

// Save implements MongoRepositorer Save
func (mr *MongoRepository) Save(model interface{}) error {
	err := mr.collection.Insert(model)
	return err
}

// GetByID implements MongoRepositorer GetByID
func (mr *MongoRepository) GetByID(id string, model interface{}) error {
	return mr.collection.FindId(bson.ObjectIdHex(id)).One(model)
}

// Delete implements Delete
func (mr *MongoRepository) Delete(model interface{}) error {
	return mr.collection.Remove(model)
}

// DeleteByID implements DeleteByID
func (mr *MongoRepository) DeleteByID(id string) error {
	return mr.collection.RemoveId(bson.ObjectIdHex(id))
}

// Count implements MongoRepositorer Count
func (mr *MongoRepository) Count() (int, error) {
	return mr.collection.Count()
}
