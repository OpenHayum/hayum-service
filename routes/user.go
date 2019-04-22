package route

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/schema"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/logger"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/service"
	"log"
	"net/http"
	"strconv"
)

var schemaDecoder = schema.NewDecoder()

type userRoute struct {
	router  Router
	service service.UserService
}

func initUserRoute(router Router) {
	userRepo := repository.NewSQLUserRepository(router.GetConn())

	userService := service.NewUserService(userRepo)
	u := &userRoute{router, userService}

	u.router.POST("/user", u.createUser)
	u.router.GET("/user/:id", u.getUser)
}

func (u *userRoute) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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

func (u *userRoute) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
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
