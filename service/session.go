package service

import (
	"context"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/util"
	"time"
)

type SessionService interface {
	Save(ctx context.Context, s *models.Session) error
	GetByID(ctx context.Context, sessionID string, userID int) (*models.Session, error)
	Update(ctx context.Context, s *models.Session) error
	DeleteByID(ctx context.Context, sessionID string, userID int) error
	Delete(ctx context.Context, s *models.Session) error
}

type sessionService struct {
	repo repository.SessionRepository
}

func NewSessionService(repo repository.SessionRepository) *sessionService {
	return &sessionService{repo}
}

func (ss *sessionService) Save(ctx context.Context, s *models.Session) error {
	s.CreatedAt = time.Now()
	// TODO: do not use UUID as session token
	s.SessionID = util.GetRandID()
	// TODO: Get Expiration time in minutes from config
	s.ExpiresAt = s.CreatedAt.Add(time.Minute)

	return ss.repo.Save(ctx, s)
}

func (ss *sessionService) GetByID(ctx context.Context, sessionID string, userID int) (*models.Session, error) {
	return ss.repo.GetByID(ctx, sessionID, userID)
}

func (ss *sessionService) Update(ctx context.Context, s *models.Session) error {
	panic("implement me")
}

func (ss *sessionService) DeleteByID(ctx context.Context, sessionID string, userID int) error {
	panic("implement me")
}

func (ss *sessionService) Delete(ctx context.Context, s *models.Session) error {
	panic("implement me")
}
