package service

import (
	"time"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
	"bitbucket.org/hayum/hayum-service/util"
)

const (
	// RoleAdmin can do everything
	RoleAdmin = "ADMIN"

	// RoleModerator can act as a moderator for uploaded contents
	RoleModerator = "MODERATOR"

	// RoleUser acts as a normal user
	RoleUser = "USER"

	// RoleArtist will have privileges of an artist
	RoleArtist = "ARTIST"
)

// UserServicer exposes methods which can be perform
type UserServicer interface {
	CreateNewUser(user *models.User) error
	GetUserByID(id string, user *models.User) error
	GetUserByEmail(email string) (*models.User, error)
	Delete(user *models.User) error
}

// UserService implements all the methods in UserServicer
type UserService struct {
	repository repository.UserRepositorer
}

// NewUserService creates a new UserService
func NewUserService() *UserService {
	return &UserService{repository.NewUserRepository()}
}

// CreateNewUser creates a new User
func (s *UserService) CreateNewUser(user *models.User) error {
	var err error
	var password string

	user.Otp = util.GenerateOTP()
	user.OtpExpirationDate = time.Now().Local().Add(time.Minute * time.Duration(30))

	password, err = util.EncryptPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = password
	user.CreatedDate = time.Now()

	err = s.repository.CreateNewUser(user)

	return err
}

// GetUserByID get User by ID
func (s *UserService) GetUserByID(id string, u *models.User) error {
	return s.repository.GetUserByID(id, u)
}

// GetUserByEmail gets User by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repository.GetUserByEmail(email)
}

// Delete deletes an User
func (s *UserService) Delete(user *models.User) error {
	return s.repository.Delete(user)
}
