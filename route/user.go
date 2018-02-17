package route

import (
	"encoding/json"
	"fmt"
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
	*Router
	service  *service.UserService
	basePath string
}

func initUserRoute(r *Router, basePath string) {
	service := service.NewUserService(r.Mongo)
	route := &userRoute{Router: r, service: service, basePath: basePath}
	route.getUser()
}

func (u *userRoute) createUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	user := new(models.User)

	schemaDecoder.Decode(user, r.PostForm)
}

func (u *userRoute) getUser() {
	u.GET(util.ConstructEndpoint(u.basePath, "/user/:id"), func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		user, err := u.service.GetUserById(ps.ByName("id"))
		if err != nil {
			log.Println("Cannot get user id")
		}
		fmt.Fprintf(w, "User Id: %s!\n", ps.ByName("id"))
		json.NewEncoder(w).Encode(user)
	})
}
