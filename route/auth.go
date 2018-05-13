package route

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/service"
	"bitbucket.org/hayum/hayum-service/util"
)

type authRoute struct {
	router Router
	s      service.AuthServicer
}

func initAuthRoute(router Router) {

	a := &authRoute{router, service.NewAuthService()}

	a.router.POST("/auth/register", a.register)
	a.router.POST("/auth/login", a.login)
	a.router.POST("/auth/logout", a.logout)
}

func (a *authRoute) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, response := a.s.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.router.JSON(w, response)
}

func (a *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	username := r.FormValue("username")
	password := r.FormValue("password")

	if username == "" || password == "" {
		log.Printf("Username or Password is not provided")
		http.Error(w, "Username or Password is not provided", http.StatusBadRequest)
		return
	}

	err, session := a.s.Login(username, password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Set-Cookie", fmt.Sprintf("sesssion-id=%s", session.ID.String()))
}

func (a *authRoute) logout(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	sessionID := util.GetCookieValue(r, "session-id")

	if err := a.s.Logout(sessionID); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	a.router.JSON(w, "")
}
