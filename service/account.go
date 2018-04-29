package service

import (
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
	"gopkg.in/mgo.v2/bson"
)

// AccountServicer AccountService exposes methods
type AccountServicer interface {
	CreateNewAccount(userID string, acc *models.Account) error
	GetAccountByID(id string, acc *models.Account) error
	UpdateAccount(acc *models.Account) error
}

// AccountService implements all the methods in AccountServicer
type AccountService struct {
	repository *repository.AccountRepository
}

func NewAccountService(r *repository.MongoRepository) *AccountService {
	return &AccountService{repository.NewAccountRepository(r)}
}

func (s *AccountService) CreateNewAccount(userID string, acc *models.Account) error {
	acc.ID = bson.NewObjectId()
	acc.UserID = userID
	return s.repository.Save(acc)
}

func (s *AccountService) GetAccountByID(id string, acc *models.Account) error {
	return s.repository.GetByID(id, acc)
}

func (s *AccountService) UpdateAccount(acc *models.Account) error {
	return s.repository.Save(acc)
}
