package repository

import (
	"bitbucket.org/hayum/hayum-service/models"
	"gopkg.in/mgo.v2/bson"
)

// UserRepository embeds a Repository
type UserRepository struct {
	*Repository
}

// NewUserRepository creates a new UserRepository
func NewUserRepository(r *Repository) *UserRepository {
	return &UserRepository{r}
}

// GetUserByID gets user by Id
func (r *UserRepository) GetUserByID(id string) (*models.User, error) {
	var user *models.User
	err := r.FindId(id).One(&user)
	return user, err
}

// GetUserByEmail gets user by email
func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user *models.User
	err := r.Find(bson.M{"email": email}).One(&user)
	return user, err
}
