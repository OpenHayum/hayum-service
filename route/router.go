package route

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/db"
	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

const apiVersion1 = "/api/v1/"

// Router contains the router and mongo instance
type Router interface {
	GET(path string, handle httprouter.Handle)
	POST(path string, handle httprouter.Handle)
	Send(w http.ResponseWriter, response interface{})
	JSON(w http.ResponseWriter, response interface{})
	GetMongo() *db.Mongo
	GetRouter() *httprouter.Router
}

type hayumRouter struct {
	router   *httprouter.Router
	mongo    *db.Mongo
	basePath string
}

// NewHayumRouter creates a new Hayum router
func NewHayumRouter(mongo *db.Mongo, basePath string) Router {
	return &hayumRouter{httprouter.New(), mongo, basePath}
}

func (hr *hayumRouter) GET(path string, handle httprouter.Handle) {
	hr.router.GET(util.ConstructEndpoint(hr.basePath, path), handle)
}

func (hr *hayumRouter) POST(path string, handle httprouter.Handle) {
	hr.router.POST(util.ConstructEndpoint(hr.basePath, path), handle)
}

func (hr *hayumRouter) GetMongo() *db.Mongo {
	return hr.mongo
}

func (hr *hayumRouter) GetRouter() *httprouter.Router {
	return hr.router
}

func (hr *hayumRouter) Send(w http.ResponseWriter, response interface{}) {
	json.NewEncoder(w).Encode(response)
}

func (hr *hayumRouter) JSON(w http.ResponseWriter, response interface{}) {
	w.Header().Add("Content-Type", "application/json")
	hr.Send(w, response)
}

// NewRouter initializes all routes of the service
func NewRouter(dbURL string, dbName string) Router {

	mongo, err := db.NewMongoSession(dbURL, dbName)

	if err != nil {
		log.Panic(err.Error())
	}

	router := NewHayumRouter(mongo, util.ConstructEndpoint(apiVersion1, "/"))

	initUserRoute(router)

	return router
}
