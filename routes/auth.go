package route

import (
	"encoding/json"
	"errors"
	hyErrors "hayum/core_apis/errors"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/service"
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"
)

type authRoute struct {
	router  Router
	service service.AuthService
}

type LoginRequestBody struct {
	Identifier string
	Password   string
}

type LoginResponseBody struct {
	User    *models.User
	Session *models.Session
}

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	u := authRoute{router, authService}

	// TODO: implement web login api using cookies
	u.router.POST("/login", u.login)
	u.router.POST("/register", u.register)
	u.router.POST("/logout", u.logout)
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	body := LoginRequestBody{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	user := &models.User{}
	session, err := a.service.Login(ctx, body.Identifier, body.Password, user)
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return
	}

	logger.Log.Info(user, session)

	response := &LoginResponseBody{user, session}

	a.router.JSON(w, response)
}

func (a *authRoute) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	body := models.User{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if err := a.service.Register(ctx, &body); err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)

}

func (a *authRoute) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	reqUserID := r.Header.Get("user-id")
	sessionID := r.Header.Get("session-id")

	if sessionID == "" || reqUserID == "" {
		http.Error(w, errors.New("SessionID or UserID is missing in header").Error(), http.StatusBadRequest)
		return
	}

	userID, _ := strconv.Atoi(reqUserID)

	session := models.Session{UserID: userID, SessionID: sessionID}

	if err := a.service.Logout(ctx, &session); err != nil {
		logger.Log.Error(err)
		status := http.StatusInternalServerError
		if err == hyErrors.ErrSessionAlreadyDeleted {
			status = http.StatusNotFound
		}
		http.Error(w, err.Error(), status)
		return
	}

	w.WriteHeader(http.StatusOK)
}
