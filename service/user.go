package service

import (
	"context"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/util"
)

type UserService interface {
	Save(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByMobile(ctx context.Context, mobile string) (*models.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *userService {
	return &userService{repo}
}

func (s *userService) Save(ctx context.Context, u *models.User) error {
	password, err := util.EncryptPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = password
	return s.repo.Save(ctx, u)
}

func (s *userService) GetByID(ctx context.Context, id int) (*models.User, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *userService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.repo.GetByEmail(ctx, email)
}

func (s *userService) GetByMobile(ctx context.Context, mobile string) (*models.User, error) {
	return s.repo.GetByMobile(ctx, mobile)
}
