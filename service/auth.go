package service

import (
	"errors"
	"log"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/models/dto"
	"bitbucket.org/hayum/hayum-service/util"
)

// AuthServicer holds the AuthService contracts
type AuthServicer interface {
	Register(user *models.User) (error, *dto.AuthRegistrationResponse)
	Login(username string, password string) (error, *models.Session)
	Logout(sessionID string) error
}

// AuthService holds the UserRepository
type AuthService struct {
	userService    UserServicer
	accountService AccountServicer
	sessionService SessionServicer
}

// NewAuthService creates a new AuthService
func NewAuthService() *AuthService {
	return &AuthService{
		NewUserService(),
		NewAccountService(),
		NewSessionService(),
	}
}

// Register handles user registration
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

// Login handles user login
func (a *AuthService) Login(username string, password string) (error, *models.Session) {
	user := &models.User{}
	if err := a.userService.GetUserByUsername(username, user); err != nil {
		log.Println("Login: Unable to get user with username", username)
		return err, nil
	}

	if err := util.CompareHashAndPassword(user.Password, password); err != nil {
		log.Println("Login: Login failed for user with ID: ", user.ID.String())
		return err, nil
	}

	return a.sessionService.CreateNewSession(user.ID.String())
}

// Logout handles user logout
func (a *AuthService) Logout(sessionID string) error {
	if err := a.sessionService.DeleteSession(sessionID); err != nil {
		log.Println("Login: Unable to delete session with ID: ", sessionID)
		return err
	}

	return nil
}
