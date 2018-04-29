package repository

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"gopkg.in/mgo.v2/bson"
)

// UserRepositorer defines UserRepository methods
type UserRepositorer interface {
	MongoRepositorer
	CreateNewUser(user *models.User) error
	GetUserByID(id string, u *models.User) error
	GetUserByEmail(email string) (*models.User, error)
}

// UserRepository embeds a Repository
type UserRepository struct {
	*MongoRepository
}

// NewUserRepository creates a new UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{NewMongoRepository(config.CollectionUser)}
}

// CreateNewUser creates a new user
func (r *UserRepository) CreateNewUser(user *models.User) error {
	user.ID = bson.NewObjectId()
	err := r.Save(&user)
	return err
}

// GetUserByID gets user by Id
func (r *UserRepository) GetUserByID(id string, u *models.User) error {
	return r.GetByID(id, u)
}

// GetUserByEmail gets user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.collection.Find(bson.M{"email": email}).One(&user)
	return user, err
}
