package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"gopkg.in/mgo.v2/bson"
)

// AccountRepositorer contains AccountRepository methods
type AccountRepositorer interface {
	MongoRepositorer
	CreateNewAccount(a *models.Account) error
}

// AccountRepository embeds MongoRepository
type AccountRepository struct {
	*MongoRepository
}

// NewAccountRepository creates a new AccountRepository
func NewAccountRepository() *AccountRepository {
	return &AccountRepository{NewMongoRepository(config.CollectionAccount)}
}

// CreateNewAccount creates a new account
func (r *AccountRepository) CreateNewAccount(acc *models.Account) error {
	acc.ID = bson.NewObjectId()
	return r.Save(&acc)
}
