package route

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/repository"
	"bitbucket.org/hayum/hayum-service/service"
)

type sessionRoute struct {
	router  Router
	service service.SessionServicer
}

func initSessionRoute(router Router) {
	repository := repository.NewMongoRepository(router.GetMongo(), config.CollectionSession)
	sessionService := service.NewSessionService(repository)
	s := &sessionRoute{router, sessionService}

	s.router.POST("/session", s.createSession)
	s.router.POST("/session/:id", s.updateSession)
	s.router.DELETE("/session/:id", s.deleteSession)
}

func (s *sessionRoute) createSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	userID := r.Header.Get("user-id")

	if err := s.service.CreateNewSession(userID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.router.Send(w, "Session created")
}

func (s *sessionRoute) updateSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sessionID := ps.ByName("id")
	log.Printf("Updating session with ID: %s", sessionID)

	if err := s.service.UpdateSession(sessionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.router.Send(w, fmt.Sprintf("Session updated with ID: %s", sessionID))
}

func (s *sessionRoute) deleteSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sessionID := ps.ByName("id")
	log.Printf("Deleting session with ID: %s", sessionID)

	if err := s.service.DeleteSession(sessionID); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	s.router.Send(w, fmt.Sprintf("Deleted session with ID: %s", sessionID))
}
