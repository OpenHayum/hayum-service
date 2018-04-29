package repository

import "bitbucket.org/hayum/hayum-service/config"

type SessionRepositorer interface {
	MongoRepositorer
}

// SessionRepository repository holds a MongoRepository
type SessionRepository struct {
	*MongoRepository
}

// NewSessionRepository creates a new SessionRepository
func NewSessionRepository() *SessionRepository {
	return &SessionRepository{NewMongoRepository(config.CollectionSession)}
}
