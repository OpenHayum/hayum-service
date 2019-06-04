package service

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"

	"github.com/pkg/errors"
)

type AuthService interface {
	Login(ctx context.Context, identifier string, password string, user *models.User) error
	Register(context.Context, *models.User) error
}

type authService struct {
	userService    UserService
	accountService AccountService
}

func NewAuthService(conn *db.Conn) *authService {
	userRepo := repository.NewSQLUserRepository(conn)
	userService := NewUserService(userRepo)
	accountRepo := repository.NewSQLAccountRepository(conn)
	acctService := NewAccountService(accountRepo)

	return &authService{userService, acctService}
}

func (a *authService) Login(ctx context.Context, identifier string, password string, user *models.User) error {
	emailAuth := newEmailAuthenticater(a.userService, ctx)
	mobileAuth := newMobileAuthenticater(a.userService, ctx)

	var hyAuth HayumAuthenticater
	hyAuth = newHayumAuthenticater(emailAuth)
	hyAuth.Add(mobileAuth)

	if !hyAuth.Authenticate(identifier, password, user) {
		logger.Log.Error("Failed to authenticate")
		err := errors.New("Failed to authenticate")
		return err
	}

	logger.Log.Infof("Successfully authenticated with identifier:%s, userId: %d", identifier, user.Id)
	return nil
}

func (a *authService) Register(ctx context.Context, user *models.User) error {
	err := a.userService.Save(ctx, user)
	if err != nil {
		return err
	}

	return a.accountService.Save(ctx, new(models.Account))
}
