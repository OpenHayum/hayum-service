package repository

import (
	"bitbucket.org/hayum/hayum-service/models"
)

type UserRepository struct {
	*Repository
}

func NewUserRepository(r *Repository) *UserRepository {
	return &UserRepository{r}
}

func (repository *UserRepository) GetUserById(id string) (*models.User, error) {
	var user *models.User
	err := repository.FindId(id).One(&user)
	return user, err
}
