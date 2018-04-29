package service

import (
	"errors"
	"log"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/models/dto"
)

// AuthServicer holds the AuthService contracts
type AuthServicer interface {
	Register(user *models.User) (error, *dto.AuthRegistrationResponse)
	// Login(username string, password string) error
	// Logout(sesssion *models.Session)
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
} //

func (a *AuthService) Register(user *models.User) (error, *dto.AuthRegistrationResponse) {
	acc := &models.Account{}
	err := a.userService.CreateNewUser(user)

	if err != nil {
		log.Println("Unable to create new User")
		return errors.New("Unable to create new User"), nil
	}

	err = a.accountService.CreateNewAccount(user.ID.String(), acc)

	if err != nil {
		log.Printf("Unable to create a new Account with UserID: %s Rollback...", user.ID)
		if err := a.userService.Delete(user); err != nil {
			log.Println("Unable to delete User. Rollback failed")
		}
		return errors.New("Unable to create new Account"), nil
	}

	return nil, &dto.AuthRegistrationResponse{user.ID.String(), acc.ID.String()}
}
