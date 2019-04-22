package route

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/service"
	"log"
	"net/http"
	"strconv"
)

type authRoute struct {
	router  Router
	service service.AuthService
}

func initAuthRoute(router Router) {
	authService := service.NewAuthService(router.GetConn())
	u := &authRoute{router, authService}

	u.router.POST("/login", u.register)
	u.router.GET("/register", u.login)
}

func (u *authRoute) register(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var user models.User

	ctx := r.Context()

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	//if u, _ := u.service.GetUserByEmail(user.Email); u != nil {
	//	http.Error(w, "User already exists!", http.StatusConflict)
	//	return
	//}

	if err := u.service.Save(ctx, &user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Header().Set("Content-Type", "application/json")
	u.router.JSON(w, user)
}

func (u *authRoute) login(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	ctx := r.Context()
	id, err := strconv.Atoi(ps.ByName("id"))

	if err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	user, err := u.service.GetByID(ctx, id)

	if err != nil {
		log.Println(err)
		log.Println("Cannot get user by id:", id)
		http.Error(w, fmt.Sprintf("Cannot get user by id: %d", id), http.StatusNotFound)
		return
	}

	u.router.JSON(w, user)
}
