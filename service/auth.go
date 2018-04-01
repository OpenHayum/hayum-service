package service

import (
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
)

// AuthServicer holds the AuthService contracts
type AuthServicer interface {
	Login(username string, password string) error
	Logout(sesssion *models.Session)
}

// AuthService holds the UserRepository
type AuthService struct {
	repository *repository.UserRepository
}
