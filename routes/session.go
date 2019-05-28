package route

import (
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/service"
	"net/http"
	"strconv"
)

type sessionRoute struct {
	router  Router
	service service.SessionService
}

func initSessionRoute(router Router) {
	sessionRepo := repository.NewSQLSessionRepository(router.GetConn())

	sessionService := service.NewSessionService(sessionRepo)
	u := &sessionRoute{router, sessionService}

	u.router.POST("/session", u.createSession)
	//u.router.GET("/session/:id", u.getSessionByID)
	//u.router.DELETE("/session/:id", u.deleteSessionByID)
}

func (s *sessionRoute) createSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	ctx := r.Context()

	userID, err := strconv.Atoi(r.Header.Get("user-id"))
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session := &models.Session{UserID: userID}

	if err := s.service.Save(ctx, session); err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	s.router.JSONWithStatus(w, http.StatusCreated, session)
}

func (s *sessionRoute) getSession(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")

}
