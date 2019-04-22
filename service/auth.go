package service

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
)

type AuthService interface {
	Login(ctx context.Context, identifier string, password string) (*models.Session, error)
	Register(ctx context.Context, user *models.User) error
}

type authService struct {
	userService    UserService
	sessionService SessionService
}

func NewAuthService(conn *db.Conn) *authService {
	sessionRepo := repository.NewSQLSessionRepository(conn)
	sessionService := NewSessionService(sessionRepo)

	userRepo := repository.NewSQLUserRepository(conn)
	userService := NewUserService(userRepo)

	return &authService{userService, sessionService}
}

func (a *authService) Login(ctx context.Context, identifier string, password string) (*models.Session, error) {
	emailAuth := newEmailAuthenticater(a.userService, ctx)
	mobileAuth := newMobileAuthenticater(a.userService, ctx)

	var hyAuth HayumAuthenticater
	hyAuth = newHayumAuthenticater(emailAuth)
	hyAuth.Add(mobileAuth)

	if hyAuth.Authenticate(identifier, password) {
		logger.Log.Info("Successfully authenticated with identifier:", identifier)
	}

	session := &models.Session{}
	if err := a.sessionService.Save(ctx, session); err != nil {
		logger.Log.Error(err)
		return nil, err
	}
	return session, nil
}

func (a *authService) Register() error {
	panic("Implement me")
}
