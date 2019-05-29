// +build integration

package test

import (
	"bytes"
	"context"
	"encoding/json"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

var (
	user *models.User
)

type hayumSuite struct {
	suite.Suite
	Conn *db.Conn
	Cfg  *viper.Viper
	ts   *httptest.Server
	URL  func(string) string
}

func (s *hayumSuite) SetupSuite() {
	s.Conn = &db.Conn{DB: db.OpenContext(context.Background(), s.Cfg)}
	s.ts = newServer(s.Conn)
	s.URL = func(pathname string) string { return s.ts.URL + "/api/v1/" + pathname }
	user = &models.User{
		Email:     "dev@gmail.com",
		FirstName: "Devajit",
		LastName:  "Asem",
		Mobile:    "6724986233",
		Password:  "7hund3r",
	}
}

func (s *hayumSuite) TearDownSuite() {
	defer s.Conn.Close()
	logger.Log.Info("Tearing down")
	dropTables := `
		SET foreign_key_checks = 0;
		DROP TABLE IF EXISTS User;
		DROP TABLE IF EXISTS Account;
		DROP TABLE IF EXISTS Follower;
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

func (s *hayumSuite) checkError(err error) {
	if err != nil {
		s.T().Error(err)
	}
}
func (s *hayumSuite) TestCreateUser() {
	ts := s.ts

	reqBody, err := json.Marshal(user)
	s.checkError(err)

	resp, err := ts.Client().Post(s.URL("user"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	assert.True(s.T(), len(respBody) > 0)
	assert.Equal(s.T(), resp.StatusCode, http.StatusCreated)
}

func (s *hayumSuite) TestGetUser() {
	resp, err := s.ts.Client().Get(s.URL("user/1"))
	s.checkError(err)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("%s", respBody)
	}

	var u models.User
	err = json.Unmarshal(respBody, &u)
	s.checkError(err)
	assert.True(s.T(), u.Email == user.Email)
}
