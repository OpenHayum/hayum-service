package repository

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/models"
)

type AccountRepository interface {
	Save(ctx context.Context, a *models.Account) error
	GetByID(ctx context.Context, id int) (*models.Account, error)
}

type sqlAccountRepo struct {
	conn *db.Conn
}

func NewSQLAccountRepository(conn *db.Conn) *sqlAccountRepo {
	return &sqlAccountRepo{conn}
}

func (r *sqlAccountRepo) Save(ctx context.Context, acc *models.Account) error {
	stmt := "INSERT INTO Account (UserId, IsPremium, AccountImageLink, CoverImageLink) VALUES (?, ?, ?, ?)"
	_, err := r.conn.ExecContext(ctx, stmt, acc.UserId, acc.IsPremium, acc.AccountImageLink, acc.CoverImageLink)
	if err != nil {
		return err
	}

	return nil
}

func (r *sqlAccountRepo) GetByID(ctx context.Context, id int) (*models.Account, error) {
	stmt := "SELECT * FROM Account WHERE Id=?"
	acc := models.Account{}
	err := r.conn.GetContext(ctx, &acc, stmt, id)
	return &acc, err
}
