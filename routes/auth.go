package route

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/service"
	"net/http"
)

type authRoute struct {
	router  Router
	service service.AuthService
}

type loginRequestBody struct {
	Identifier string
	Password   string
}

type loginResponseBody struct {
	User    *models.User
	Session *models.Session
}

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	u := authRoute{router, authService}

	// TODO: implement web login api using cookies
	u.router.POST("/login", u.login)

	u.router.GET("/register", u.register)
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	body := loginRequestBody{}

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

	response := &loginResponseBody{user, session}

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
