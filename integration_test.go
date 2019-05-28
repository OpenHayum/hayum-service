// +build integration

package main

import (
	"context"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"testing"

	"github.com/spf13/viper"
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

func (s *hayumSuite) TeardownSuite() {
	s.Conn.Close()
}

func TestHayumSuite(t *testing.T) {
	logger.Init()
	cfg := config.New()
	hySuite := &hayumSuite{
		Cfg: cfg,
	}

	suite.Run(t, hySuite)
}
