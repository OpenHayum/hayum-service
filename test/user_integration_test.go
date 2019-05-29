package test

import (
	"encoding/json"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

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
