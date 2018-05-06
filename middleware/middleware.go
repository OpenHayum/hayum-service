package middleware

import (
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/models"

	"bitbucket.org/hayum/hayum-service/service"
)

var sessionService *service.SessionService

// InitMiddlewareServices initialize all middleware services only once
func InitMiddlewareServices() {
	if sessionService == nil {
		sessionService = service.NewSessionService()
	}
}

// Authorize check for session
// TODO: manage a .yaml file for routes that requires authentication
// such as Artist visiting his/her profile should show a different view than
// a normal user visiting someone else'e profile
// TODO: shift session management to redis
func Authorize(rw http.ResponseWriter, r *http.Request) {
	sessionIDCookie, err := r.Cookie("session-id")
	userIDCookie, _ := r.Cookie("user-id")
	if err != nil {
		log.Println("Authorize: Unable to read cookie", err)
	}
	sessionID := sessionIDCookie.Value
	userID := userIDCookie.Value
	session := new(models.Session)

	if err := sessionService.GetSessionByID(sessionID, session); err != nil {
		log.Printf("Authorize: Unable to get session by ID: %s", sessionID)
	}

	// TODO: handle redirects here
	if session.UserID == userID {
		log.Printf("Authorize: Authorized user %s", userID)
	}
	log.Printf("Authorize: Unauthorized user %s", userID)
}
