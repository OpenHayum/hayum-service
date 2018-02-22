package route

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/schema"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/repository"
	"bitbucket.org/hayum/hayum-service/util"

	"bitbucket.org/hayum/hayum-service/service"
	"github.com/julienschmidt/httprouter"
)

var schemaDecoder = schema.NewDecoder()

type userRoute struct {
	router  *Router
	service *service.UserService
}

func initUserRoute(router *Router, basePath string) {
	service := service.NewUserService(repository.NewRepository(router.Mongo, config.CollectionUser))
	u := &userRoute{router: router, service: service}

	u.router.POST(util.ConstructEndpoint(basePath, "/user"), u.createUser)
	u.router.GET(util.ConstructEndpoint(basePath, "/user/:id"), u.getUser)
}

func (u *userRoute) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if u, _ := u.service.GetUserByEmail(user.Email); u != nil {
		http.Error(w, "User already exists!", http.StatusConflict)
		return
	}

	if err := u.service.CreateNewUser(&user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	u.router.Send(w, user)
}

func (u *userRoute) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := u.service.GetUserByID(ps.ByName("id"))

	if err != nil {
		log.Println("Cannot get user id")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	u.router.Send(w, user)
}
