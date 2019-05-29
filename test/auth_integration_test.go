package test

import (
	"bytes"
	"encoding/json"
	"hayum/core_apis/logger"
	"hayum/core_apis/routes"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

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
