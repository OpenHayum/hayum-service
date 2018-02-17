package route

import (
	"fmt"
	"net/http"

	"bitbucket.org/hayum/hayum-service/service"
	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

type UserRoute struct{}

func (r *Router) UserRoute(basePath string) {
	userService := service.NewUserService(r.Mongo)
	r.GET(util.ConstructEndpoint(basePath, "/:name"), getUser(use))
}
