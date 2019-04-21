package service

import (
	"context"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/util"
	"time"
)

type UserService interface {
	Create(ctx context.Context, u *models.User) error
	FindByID(ctx context.Context, id int) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) Create(ctx context.Context, u *models.User) error {
	u.CreatedDate = time.Now()
	password, err := util.EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return s.repo.Save(ctx, u)
}

func (s *userService) FindByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}
