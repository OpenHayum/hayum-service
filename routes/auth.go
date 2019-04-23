package route

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/logger"
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

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	u := authRoute{router, authService}

	u.router.POST("/login", u.login)
	u.router.GET("/register", u.register)
}

func (a *authRoute) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	panic("Implement me")
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()

	body := loginRequestBody{}

	if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	session, err := a.service.Login(ctx, body.Identifier, body.Password)
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, "Not Authorized", http.StatusForbidden)
		return
	}

	a.router.JSON(w, session)
}