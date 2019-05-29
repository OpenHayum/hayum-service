package route

import (
	"encoding/json"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/service"
	"net/http"

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
