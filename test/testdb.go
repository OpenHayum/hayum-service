package test

import (
	"context"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
)

func truncate(conn *db.Conn) {
	stmt := `
		SET foreign_key_checks=0;
		TRUNCATE TABLE User;
		TRUNCATE TABLE Session;
		SET foreign_key_checks=1;
	`
	if _, err := conn.Exec(stmt); err != nil {
		logger.Log.Errorf("error truncating database tables: %v", err)
	}
}

// persist user in db
func seedUser(u *models.User, conn *db.Conn) {
	userRepo := repository.NewSQLUserRepository(conn)
	if err := userRepo.Save(context.Background(), u); err != nil {
		logger.Log.Error(err)
	}
}

// persist session in db
func seedSession(u *models.User, conn *db.Conn) {
	repo := repository.NewSQLSessionRepository(conn)
	if err := repo.Save(context.Background(), &models.Session{UserID: u.Id}); err != nil {
		logger.Log.Error(err)
	}
}
