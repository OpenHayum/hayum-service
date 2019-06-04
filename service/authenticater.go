package service

import (
	"context"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/util"
)

// Authenticater interface to implement different type of auth using different columns
type Authenticater interface {
	// identifier can be one of [`Email`, `Mobile`]
	// TODO: support username
	authenticate(identifier string, password string, user *models.User) bool
}

type HayumAuthenticater interface {
	Add(a Authenticater)
	Authenticate(identifier string, password string, user *models.User) bool
}

type hayumAuthenticater struct {
	auth Authenticater
	next *hayumAuthenticater
}

func newHayumAuthenticater(a Authenticater) *hayumAuthenticater {
	return &hayumAuthenticater{a, nil}
}

func (ha *hayumAuthenticater) Add(a Authenticater) {
	ha.next = &hayumAuthenticater{a, nil}
}

func (ha *hayumAuthenticater) Authenticate(identifier string, password string, user *models.User) bool {
	curr := ha

	for curr != nil {
		isAuthenticated := curr.auth.authenticate(identifier, password, user)
		if isAuthenticated {
			return true
		}
		curr = curr.next
	}

	return false
}

type emailAuthenticater struct {
	s   UserService
	ctx context.Context
}

type mobileAuthenticater struct {
	s   UserService
	ctx context.Context
}

type usernameAuthenticater struct {
	s   UserService
	ctx context.Context
}

func newEmailAuthenticater(s UserService, ctx context.Context) *emailAuthenticater {
	return &emailAuthenticater{s, ctx}
}

func isPasswordEqual(passwordFromInput string, storedHashedPassword string) bool {
	return util.CompareHashAndPassword(storedHashedPassword, passwordFromInput) == nil
}

func (e *emailAuthenticater) authenticate(email string, password string, user *models.User) bool {
	u, err := e.s.GetByEmail(e.ctx, email)
	if err != nil || (*user == models.User{}) {
		logger.Log.Error(err)
		return false
	}
	*user = *u

	return isPasswordEqual(password, user.Password)
}

func newMobileAuthenticater(s UserService, ctx context.Context) *mobileAuthenticater {
	return &mobileAuthenticater{s, ctx}
}

func (e *mobileAuthenticater) authenticate(mobile string, password string, user *models.User) bool {
	u, err := e.s.GetByMobile(e.ctx, mobile)
	if err != nil || (*user == models.User{}) {
		logger.Log.Error(err)
		return false
	}
	*user = *u

	return isPasswordEqual(password, user.Password)
}

func newUsernameAuthenticater(s UserService, ctx context.Context) *usernameAuthenticater {
	return &usernameAuthenticater{s, ctx}
}
