package service

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
)

type userRepository interface {
	CreateNewUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(mongo *config.Mongo) *UserService {
	return &UserService{repository.NewUserRepository(repository.NewRepository(mongo, "user"))}
}

func (s *UserService) CreateNewUser(user *models.User) error {
	err := s.repository.CreateNewUser(user)
	return err
}

func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repository.GetUserByID(id)
}

func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repository.GetUserByEmail(email)
}
