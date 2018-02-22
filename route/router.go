package route

import (
	"encoding/json"
	"log"
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

// NewRouter initializes all routes of the service
func NewRouter(dbURL string, dbName string) *Router {

	mongo, err := config.NewMongoSession(dbURL, dbName)

	if err != nil {
		log.Panic(err.Error())
	}

	router := Router{Router: httprouter.New(), Mongo: mongo}

	initUserRoute(&router, util.ConstructEndpoint(apiVersion1, "/"))

	return &router
}

// Send writes the model to ResponseWriter
func (r *Router) Send(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}
