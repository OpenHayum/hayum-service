package repository

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/models"
)

type UserRepository interface {
	Save(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id int) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByMobile(ctx context.Context, mobile string) (*models.User, error)
}

type sqlUserRepo struct {
	conn *db.Conn
}

func NewSQLUserRepository(conn *db.Conn) *sqlUserRepo {
	return &sqlUserRepo{conn}
}

func (r *sqlUserRepo) Save(ctx context.Context, u *models.User) error {
	stmt := "INSERT INTO User (FirstName, LastName, Email, Mobile, Password, CreatedDate) VALUES (?, ?, ?, ?, ?, ?)"
	_, err := r.conn.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Mobile, u.Password, u.CreatedDate)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlUserRepo) GetByID(ctx context.Context, id int) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Id=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, id)
	return &user, err
}

func (r *sqlUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Email=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, email)
	return &user, err
}

func (r *sqlUserRepo) GetByMobile(ctx context.Context, mobile string) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Mobile=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, mobile)
	return &user, err
}
