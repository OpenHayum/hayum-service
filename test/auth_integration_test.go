package test

import (
	"bytes"
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"hayum/core_apis/models"
	"hayum/core_apis/routes"
	"io/ioutil"
	"net/http"
	"strconv"
)

func (s *hayumSuite) loginUser(user *models.User) *route.LoginResponseBody {
	s.createUser(user)
	reqBody, err := json.Marshal(&route.LoginRequestBody{Identifier: user.Email, Password: user.Password})
	s.checkError(err)

	resp, err := s.ts.Client().Post(s.URL("login"), "application/json", bytes.NewReader(reqBody))
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("Failed to login %s", resp.Status)
	}

	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	s.checkError(err)

	var loginResp route.LoginResponseBody
	err = json.Unmarshal(respBody, &loginResp)
	s.checkError(err)

	return &loginResp
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
	loginResp := s.loginUser(user)
	assert.True(s.T(), loginResp.User.Email == user.Email, "Emails should be equal")

	truncate(s.Conn)
}

func (s *hayumSuite) TestLogout() {
	user := getUser()
	loginResp := s.loginUser(user)

	req, _ := http.NewRequest("POST", s.URL("logout"), nil)
	req.Header.Add("user-id", strconv.Itoa(loginResp.User.Id))
	req.Header.Add("session-id", loginResp.Session.SessionID)

	resp, err := s.ts.Client().Do(req)
	s.checkError(err)

	if resp.StatusCode != http.StatusOK {
		s.T().Fatalf("%s %s", err, resp.Status)
	}

	resp, err = s.ts.Client().Do(req)
	s.checkError(err)

	assert.Equal(s.T(), http.StatusNotFound, resp.StatusCode, "Should return 404(Session already deleted)")

	truncate(s.Conn)
}
