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
	Login(ctx context.Context, identifier string, password string, user *models.User) (*models.Session, error)
	Register(context.Context, *models.User) error
	Logout(context.Context, *models.Session) error
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

func (a *authService) Login(ctx context.Context, identifier string, password string, user *models.User) (*models.Session, error) {
	emailAuth := newEmailAuthenticater(a.userService, ctx)
	mobileAuth := newMobileAuthenticater(a.userService, ctx)

	var hyAuth HayumAuthenticater
	hyAuth = newHayumAuthenticater(emailAuth)
	hyAuth.Add(mobileAuth)

	if !hyAuth.Authenticate(identifier, password, user) {
		logger.Log.Error("Failed to authenticate")
		err := errors.New("Failed to authenticate")
		return nil, err
	}

	logger.Log.Infof("Successfully authenticated with identifier:%s, userId: %d", identifier, user.Id)

	session := &models.Session{UserID: user.Id}
	if err := a.sessionService.Save(ctx, session); err != nil {
		logger.Log.Error(err)
		return nil, err
	}

	return session, nil
}

func (a *authService) Register(ctx context.Context, user *models.User) error {
	return a.userService.Save(ctx, user)
}

func (a *authService) Logout(ctx context.Context, session *models.Session) error {
	return a.sessionService.DeleteByID(ctx, session.SessionID, session.UserID)
}
