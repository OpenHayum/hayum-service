package repository

// SessionRepository repository holds a MongoRepository
type SessionRepository struct {
	*MongoRepository
}

// NewSessionRepository creates a new SessionRepository
func NewSessionRepository(r *MongoRepository) *SessionRepository {
	return &SessionRepository{r}
}
