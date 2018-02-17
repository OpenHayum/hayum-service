package service

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
)

type userRepository interface {
	GetUserById(id string) (*models.User, error)
}

type UserService struct {
	repository userRepository
}

func NewUserService(mongo *config.Mongo) *UserService {
	return &UserService{repository.NewUserRepository(repository.NewRepository(mongo, "user"))}
}

func (s *UserService) GetUserById(id string) (*models.User, error) {
	return s.repository.GetUserById(id)
}
