// +build integration

package main

import (
	"context"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/service"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type hayumSuite struct {
	suite.Suite
	Conn *db.Conn
	Cfg  *viper.Viper
}

func (s *hayumSuite) SetupSuite() {
	s.Conn = &db.Conn{db.OpenContext(context.Background(), s.Cfg)}
}

func (s *hayumSuite) TearDownSuite() {
	defer s.Conn.Close()
	logger.Log.Info("Tearing down")
	dropTables := `
		SET foreign_key_checks = 0;
		-- DROP TABLE IF EXISTS User;
		-- DROP TABLE IF EXISTS Account;
		-- DROP TABLE IF EXISTS Follower;
		SET foreign_key_checks = 1;
	`
	result, err := s.Conn.Exec(dropTables)
	logger.Log.Info(result)
	if err != nil {
		logger.Log.Fatal(err)
	}

}

func TestHayumSuite(t *testing.T) {
	logger.Init()
	cfg := config.New()
	hySuite := &hayumSuite{
		Cfg: cfg,
	}

	suite.Run(t, hySuite)
}

func (s *hayumSuite) TestCreateUser() {
	user := &models.User{
		Email:     "dev@gmail.com",
		FirstName: "Devajit",
		LastName:  "Asem",
		Mobile:    "6724986233",
		Password:  "7hund3r",
	}

	userRepo := repository.NewSQLUserRepository(s.Conn)
	userService := service.NewUserService(userRepo)
	err := userService.Save(context.Background(), user)

	assert.Equal(s.T(), nil, err)
}

func (s *hayumSuite) TestGetUser() {
	// userRepo := repository.NewSQLUserRepository(s.Conn)
	// userService := service.NewUserService(userRepo)
	// userService.GetByEmail()
}
