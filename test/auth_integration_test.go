package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"hayum/core_apis/models"
	"hayum/core_apis/route"
	"io/ioutil"
	"net/http"
)

func (s *hayumSuite) loginUser(user *models.User) *http.Response {
	s.createUser(user)
	reqBody, err := json.Marshal(&route.LoginRequestBody{Identifier: user.Email, Password: user.Password})
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("login"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("Failed to login %s", resp.Status)
	}

	if resp.Header.Get("Set-Cookie") == "" {
		s.T().Fatalf("Cookie not set in response header!")
	}

	return resp
}

func (s *hayumSuite) getLoginResponse(resp *http.Response) *models.User {
	defer resp.Body.Close()
	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	var userRes models.User
	err = json.Unmarshal(respBody, &userRes)
	s.checkError(err)

	return &userRes
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
	loginResp := s.getLoginResponse(resp)
	assert.True(s.T(), loginResp.Email == user.Email, "Emails should be equal")

	truncate(s.Conn)
}

func (s *hayumSuite) TestLogout() {
	user := getUser()

	resp := s.loginUser(user)
	req, _ := http.NewRequest("POST", s.URL("logout"), nil)
	req.Header.Add("Cookie", resp.Header.Get("Set-Cookie"))
	resp, err := s.ts.Client().Do(req)
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("%s %s", err, resp.Status)
	}

	req.Header.Set("Cookie", "")

	resp, err = s.ts.Client().Do(req)
	s.checkError(err)

	assert.Equal(s.T(), http.StatusBadRequest, resp.StatusCode, "Should return 400")
	truncate(s.Conn)
}
