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

func getCookieValue(r *http.Request, name string) string {
	cookie, err := r.Cookie(name)
	if err != nil {
		log.Println("getCookieValue: Unable to read cookie with name: ", name, err)
		return ""
	}
	return cookie.Value
}

// Authorize check for session
// TODO: manage a .yaml file for routes that requires authentication
// such as Artist visiting his/her profile should show a different view than
// a normal user visiting someone else'e profile
// TODO: shift session management to redis
func Authorize(rw http.ResponseWriter, r *http.Request) {
	sessionID := getCookieValue(r, "session-id")
	userID := getCookieValue(r, "user-id")
	session := new(models.Session)

	if sessionID == "" {
		// TODO: if there is no `session-id` present in cookie
		// check if route requires authentication and redirect to login
		// else just let it through without blocking
	}

	// TODO: check for expired sessions, delete session if it is expired
	// and redirect to login

	if err := sessionService.GetSessionByID(sessionID, session); err != nil {
		log.Printf("Authorize: Unable to get session by ID: %s", sessionID)
	}

	// TODO: handle redirects here
	if session.UserID == userID {
		log.Printf("Authorize: Authorized user %s", userID)
	}
	log.Printf("Authorize: Unauthorized user %s", userID)
}
