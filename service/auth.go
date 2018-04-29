package service

import (
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/models/dto"
)

// AuthServicer holds the AuthService contracts
type AuthServicer interface {
	Register(user *models.User) (error, *dto.AuthRegistrationResponse)
	Login(username string, password string) error
	Logout(sesssion *models.Session)
}

// AuthService holds the UserRepository
type AuthService struct {
	userService    UserServicer
	accountService AccountServicer
}

func NewAuthService() *AuthService {
	return &AuthService{
		NewUserService(),
		NewAccountService(),
	}
}

func (a *AuthService) Register(user *models.User) *dto.AuthRegistrationResponse {
	acc := &models.Account{}
	a.userService.CreateNewUser(user)
	a.accountService.CreateNewAccount(user.ID.String(), acc)

	return &dto.AuthRegistrationResponse{user.ID.String(), acc.ID.String()}
}
