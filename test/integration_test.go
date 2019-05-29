package test

import (
	"bytes"
	"context"
	"encoding/json"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/routes"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
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
		s.T().Fatal(err)
	}
}

func (s *hayumSuite) createUser(user *models.User) *http.Response {
	reqBody, err := json.Marshal(user)
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("user"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)
	return resp
}

// ************************************* User ****************************************

func (s *hayumSuite) TestCreateUser() {
	user := getUser()
	resp := s.createUser(user)
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	assert.True(s.T(), len(respBody) > 0)
	assert.Equal(s.T(), resp.StatusCode, http.StatusCreated)
	truncate(s.Conn)
}

func (s *hayumSuite) TestGetUser() {
	user := getUser()
	seedUser(user, s.Conn)
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
	logger.Log.Info(u)
	assert.True(s.T(), u.Email == user.Email)
	truncate(s.Conn)
}

// ************************************* Session ****************************************

func (s *hayumSuite) TestCreateSession() {
	user := getUser()
	seedUser(user, s.Conn)
	req, _ := http.NewRequest("POST", s.URL("session"), nil)
	req.Header.Add("user-id", "1")

	resp, err := s.ts.Client().Do(req)
	s.checkError(err)

	if resp.StatusCode != http.StatusCreated {
		s.T().Fatalf("%s %s", err, resp.Status)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	var session models.Session
	err = json.Unmarshal(respBody, &session)
	s.checkError(err)

	assert.True(s.T(), session.UserID == 1)
	truncate(s.Conn)
}

// ************************************* Auth(Register/Login) ****************************************

func (s *hayumSuite) TestRegister() {
	user := getUser()
	reqBody, err := json.Marshal(user)
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("register"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)

	if resp.StatusCode != http.StatusCreated {
		s.T().Fatalf("%v %s", err, resp.Status)
	}

	truncate(s.Conn)
}

func (s *hayumSuite) TestLogin() {
	user := getUser()
	s.createUser(user)
	reqBody, err := json.Marshal(&route.LoginRequestBody{Identifier: user.Email, Password: user.Password})
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("login"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("%v %s", err, resp.Status)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	var loginResp route.LoginResponseBody
	err = json.Unmarshal(respBody, &loginResp)
	s.checkError(err)
	logger.Log.Info(loginResp)
	assert.True(s.T(), loginResp.User.Email == user.Email, "Emails should be equal")

	truncate(s.Conn)
}
