package route

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

const apiVersion1 = "/api/v1/"

// Router contains the router and mongo instance
type Router struct {
	*httprouter.Router
	*config.Mongo
}

// Init initializes all routes of the service
func (r *Router) Init() {
	initUserRoute(r, util.ConstructEndpoint(apiVersion1, "/"))
}

// Send writes the model to ResponseWriter
func (r *Router) Send(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
