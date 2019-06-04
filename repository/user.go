package repository

import (
	"context"
	"database/sql"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
)

type UserRepository interface {
	Save(ctx context.Context, u *models.User) error
	GetByID(ctx context.Context, id int64) (*models.User, error)
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	GetByMobile(ctx context.Context, mobile string) (*models.User, error)
	GetByMobileOrEmail(ctx context.Context, mobile, email string) (*models.User, error)
}

type sqlUserRepo struct {
	conn *db.Conn
}

func NewSQLUserRepository(conn *db.Conn) *sqlUserRepo {
	return &sqlUserRepo{conn}
}

func (r *sqlUserRepo) Save(ctx context.Context, u *models.User) error {
	stmt := "INSERT INTO User (FirstName, LastName, Email, Mobile, Password, CreatedDate) VALUES (?, ?, ?, ?, ?, ?)"
	res, err := r.conn.ExecContext(ctx, stmt, u.FirstName, u.LastName, u.Email, u.Mobile, u.Password, u.CreatedDate)
	if err != nil {
		return err
	}

	u.Id, err = res.LastInsertId()
	return err
}

func (r *sqlUserRepo) GetByID(ctx context.Context, id int64) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Id=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, id)
	if err == sql.ErrNoRows {
		return &user, nil
	}
	return &user, err
}

func (r *sqlUserRepo) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Email=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, email)
	logger.Log.Info("User:", user)
	if err == sql.ErrNoRows {
		return &user, nil
	}
	return &user, err
}

func (r *sqlUserRepo) GetByMobile(ctx context.Context, mobile string) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Mobile=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, mobile)
	logger.Log.Info("User:", user)
	if err == sql.ErrNoRows {
		return &user, nil
	}
	return &user, err
}

func (r *sqlUserRepo) GetByMobileOrEmail(ctx context.Context, mobile, email string) (*models.User, error) {
	stmt := "SELECT * FROM User WHERE Mobile=? OR Email=?"
	user := models.User{}
	err := r.conn.GetContext(ctx, &user, stmt, mobile, email)
	if err == sql.ErrNoRows {
		return &user, nil
	}
	return &user, err
}
