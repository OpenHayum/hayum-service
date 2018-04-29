package service

import (
	"log"
	"strconv"
	"time"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
	"gopkg.in/mgo.v2/bson"
)

// SessionServicer holds the SessionService contracts
type SessionServicer interface {
	CreateNewSession(userID string) error
	DeleteSession(sessionID string) error
	UpdateSession(sessionID string) error
}

// SessionService holds the SessionRepository
type SessionService struct {
	repository         repository.SessionRepositorer
	expirationTimeInMS int
}

// NewSessionService creates a new SessionService
func NewSessionService() *SessionService {
	expirationTimeInMS, err := strconv.Atoi(config.App.GetString("sessionExpirationTimeInMS"))
	if err != nil {
		log.Panic("Unable to convert string to int", err)
	}
	return &SessionService{repository.NewSessionRepository(), expirationTimeInMS}
}

// CreateNewSession creates a new Session
func (s *SessionService) CreateNewSession(userID string) error {
	session := new(models.Session)
	session.ID = bson.NewObjectId()
	session.UserID = userID
	session.ExpiresAt = time.Now().Local().Add(time.Millisecond * time.Duration(s.expirationTimeInMS))
	return s.repository.Save(session)
}

// DeleteSession deletes a Session
func (s *SessionService) DeleteSession(sessionID string) error {
	return s.repository.DeleteByID(sessionID)
}

// UpdateSession updates a session
func (s *SessionService) UpdateSession(sessionID string) error {
	session := new(models.Session)
	if err := s.repository.GetByID(sessionID, session); err != nil {
		log.Fatalf("Session with sessionID: %s not found.", sessionID)
		return err
	}
	session.ExpiresAt = time.Now().Local().Add(time.Millisecond * time.Duration(s.expirationTimeInMS))
	return s.repository.Save(session)
}
