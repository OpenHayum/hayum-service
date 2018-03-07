package service

import (
	"time"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
	"bitbucket.org/hayum/hayum-service/util"
)

// UserServicer exposes methods which can be perform
type UserServicer interface {
	CreateNewUser(user *models.User) error
	GetUserByID(id string) (*models.User, error)
	GetUserByEmail(email string) (*models.User, error)
}

// UserService implements all the methods in UserServicer
type UserService struct {
	repository *repository.UserRepository
}

// NewUserService creates a new UserService
func NewUserService(r *repository.Repository) *UserService {
	return &UserService{repository.NewUserRepository(r)}
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

	err = s.repository.CreateNewUser(user)

	return err
}

// GetUserByID get User by ID
func (s *UserService) GetUserByID(id string) (*models.User, error) {
	return s.repository.GetUserByID(id)
}

// GetUserByEmail gets User by email
func (s *UserService) GetUserByEmail(email string) (*models.User, error) {
	return s.repository.GetUserByEmail(email)
}
