package route

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	hyErrors "hayum/core_apis/errors"
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

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	u := authRoute{router, authService}

	// TODO: implement web login api using cookies
	u.router.POST("/login", u.login)
	u.router.POST("/register", u.register)
	u.router.POST("/logout", u.logout)
}

func handleSessionErr(err error, w http.ResponseWriter) {
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, hyErrors.ErrUnableToGetSession.Error(), http.StatusInternalServerError)
		return
	}
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
	err := a.service.Login(ctx, body.Identifier, body.Password, user)
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return
	}

	session, err := db.Store.Get(r, "hayum-session")
	handleSessionErr(err, w)

	session.Values["user-id"] = user.Id
	session.Values["email"] = user.Email
	err = session.Save(r, w)
	handleSessionErr(err, w)
	//logger.Log.Info(session)

	a.router.JSON(w, user)
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
	if r.Header.Get("Cookie") == "" {
		http.Error(w, http.ErrNoCookie.Error(), http.StatusBadRequest)
		return
	}

	session, err := db.Store.Get(r, config.SessionName)
	if err != nil {
		http.Error(w, http.ErrNoCookie.Error(), http.StatusForbidden)
		return
	}

	if err := db.Store.Delete(r, w, session); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
