package test

import (
	"hayum/core_apis/models"
	"time"

	"github.com/google/uuid"
)

// generate a random user model for testing
func getUser() *models.User {
	randString := uuid.New().String()
	return &models.User{
		Email:       randString[:10] + "@gmail.com",
		FirstName:   randString[:6],
		LastName:    randString[7:15],
		Mobile:      "6724986233",
		Password:    randString[18:],
		CreatedDate: time.Now(),
	}
}
