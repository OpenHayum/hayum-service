package route

import (
	"log"
	"net/http"

	"github.com/gorilla/schema"

	"bitbucket.org/hayum/hayum-service/models"
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
	service := service.NewUserService(router.Mongo)
	u := &userRoute{router: router, service: service}

	u.router.POST(util.ConstructEndpoint(basePath, "/user"), u.createUser)
	u.router.GET(util.ConstructEndpoint(basePath, "/user/:id"), u.getUser)
}

func (u *userRoute) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	user := new(models.User)

	schemaDecoder.Decode(user, r.PostForm)
}

func (u *userRoute) getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	user, err := u.service.GetUserById(ps.ByName("id"))

	if err != nil {
		log.Println("Cannot get user id")
	}

	u.router.Send(w, user)
}
