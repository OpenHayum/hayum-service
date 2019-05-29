package repository

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/errors"
	"hayum/core_apis/models"
)

type SessionRepository interface {
	Save(ctx context.Context, s *models.Session) error
	GetByID(ctx context.Context, sessionID string, userID int) (*models.Session, error)
	Update(ctx context.Context, s *models.Session) error
	DeleteByID(ctx context.Context, sessionID string, userID int) error
}

type sqlSessionRepo struct {
	conn *db.Conn
}

func NewSQLSessionRepository(conn *db.Conn) *sqlSessionRepo {
	return &sqlSessionRepo{conn}
}

func (r *sqlSessionRepo) Save(ctx context.Context, s *models.Session) error {
	stmt := "INSERT INTO Session (SessionId, UserId, ExpiresAt, CreatedAt) VALUES (?, ?, ?, ?)"
	_, err := r.conn.ExecContext(ctx, stmt, s.SessionID, s.UserID, s.ExpiresAt, s.CreatedAt)
	return err
}

func (r *sqlSessionRepo) GetByID(ctx context.Context, sessionID string, userID int) (*models.Session, error) {
	stmt := "SELECT * FROM Session WHERE SessionId=? AND UserId=?"
	session := models.Session{}
	err := r.conn.GetContext(ctx, &session, stmt)
	return &session, err
}

func (r *sqlSessionRepo) Update(ctx context.Context, s *models.Session) error {
	panic("implement me")
}

func (r *sqlSessionRepo) DeleteByID(ctx context.Context, sessionID string, userID int) error {
	stmt := "DELETE FROM Session WHERE SessionId=? AND UserId=?"
	res, err := r.conn.ExecContext(ctx, stmt, sessionID, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAffected == 0 {
		return errors.ErrSessionAlreadyDeleted
	}
	return nil
}
