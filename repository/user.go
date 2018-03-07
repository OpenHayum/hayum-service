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
