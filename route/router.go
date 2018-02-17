package route

import (
	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

const apiVersion1 = "/api/v1/"

type Router struct {
	*httprouter.Router
	*config.Mongo
}

func (r *Router) Init() {
	r.UserRoute(util.ConstructEndpoint(apiVersion1, "/user/"))
	r.ItemRoute(util.ConstructEndpoint(apiVersion1, "/item/"))
}
