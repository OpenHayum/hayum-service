package test

import (
	"bytes"
	"encoding/json"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/routes"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/stretchr/testify/assert"
)

func (s *hayumSuite) loginUser(user *models.User) *http.Response {
	s.createUser(user)
	reqBody, err := json.Marshal(&route.LoginRequestBody{Identifier: user.Email, Password: user.Password})
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("login"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)
	return resp
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
	resp := s.loginUser(user)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("Failed to login %s", resp.Status)
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

func (s *hayumSuite) TestLogout() {
	user := getUser()
	resp := s.loginUser(user)

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	var loginResp route.LoginResponseBody
	err = json.Unmarshal(respBody, &loginResp)
	s.checkError(err)

	req, _ := http.NewRequest("POST", s.URL("logout"), nil)
	req.Header.Add("user-id", strconv.Itoa(loginResp.User.Id))
	req.Header.Add("session-id", loginResp.Session.SessionID)

	resp, err = s.ts.Client().Do(req)
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("%s %s", err, resp.Status)
	}

	truncate(s.Conn)
}
