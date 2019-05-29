package test

import (
	"encoding/json"
	"hayum/core_apis/models"
	"io/ioutil"
	"net/http"

	"github.com/stretchr/testify/assert"
)

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
