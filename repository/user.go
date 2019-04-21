package repository

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/models"
)

type UserRepository interface {
	Save(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
}

type sqlUserRepo struct {
	conn *db.Conn
}

func NewSQLUserRepository(conn *db.Conn) *sqlUserRepo {
	return &sqlUserRepo{conn}
}

func (r *sqlUserRepo) Save(ctx context.Context, u *models.User) error {
	_, err := r.conn.ExecContext(ctx, "INSERT INTO User (FirstName, LastName, Email, Mobile, Password, CreatedDate) VALUES (?, ?, ?, ?, ?, ?)",
		u.FirstName, u.LastName, u.Email, u.Mobile, u.Password, u.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlUserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, "SELECT * FROM User WHERE Id=?", id)
	return &user, err
}
