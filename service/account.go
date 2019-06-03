package service

import (
	"context"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
)

type AccountService interface {
	Save(ctx context.Context, a *models.Account) error
	GetByID(ctx context.Context, id int) (*models.Account, error)
}

type accountService struct {
	repo repository.AccountRepository
}

func NewAccountService(repo repository.AccountRepository) *accountService {
	return &accountService{repo}
}

func (s *accountService) Save(ctx context.Context, account *models.Account) error {
	return s.repo.Save(ctx, account)
}

func (s *accountService) GetByID(ctx context.Context, id int) (*models.Account, error) {
	return s.repo.GetByID(ctx, id)
}
