package route

import (
	"encoding/gob"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/errors"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/service"
	"net/http"
)

type authRoute struct {
	router  Router
	service service.AuthService
}

type LoginRequestBody struct {
	Identifier string
	Password   string
}

type SessionResponse struct {
	UserId          int64
	Email           string
	IsAuthenticated bool
}

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	gob.Register(SessionResponse{})
	u := authRoute{router, authService}

	// TODO: implement web login api using cookies
	u.router.POST("/login", u.login)
	u.router.POST("/register", u.register)
	u.router.POST("/logout", u.logout)
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	body := LoginRequestBody{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if errors.CheckAndSendResponseErrorWithStatus(err, w, http.StatusBadRequest) {
		return
	}

	user := &models.User{}
	err = a.service.Login(ctx, body.Identifier, body.Password, user)
	if errors.CheckAndSendResponseErrorWithStatus(err, w, http.StatusForbidden) {
		return
	}

	session, err := db.Store.Get(r, config.SessionName)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	sessionRes := &SessionResponse{user.Id, user.Email, true}
	session.Values["user"] = sessionRes

	err = session.Save(r, w)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	a.router.JSON(w, user)
}

func (a *authRoute) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	body := models.User{}
	err := json.NewDecoder(r.Body).Decode(&body)
	if errors.CheckAndSendResponseErrorWithStatus(err, w, http.StatusBadRequest) {
		return
	}

	err = a.service.Register(ctx, &body)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *authRoute) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	if r.Header.Get("Cookie") == "" {
		errors.CheckAndSendResponseErrorWithStatus(http.ErrNoCookie, w, http.StatusBadRequest)
	}

	session, err := db.Store.Get(r, config.SessionName)
	logger.Log.Info(session, session.Values["user"])

	// check for request which are already logged out
	if err != nil || session.Values["user"] == nil {
		if errors.CheckAndSendResponseErrorWithStatus(err, w, http.StatusForbidden) {
			return
		}
	}

	session.Options.MaxAge = -1
	session.Values["user"] = &SessionResponse{}

	err = session.Save(r, w)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	err = db.Store.Delete(r, w, session)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}
