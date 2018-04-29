package route

import (
	"encoding/json"
	"net/http"

	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

const apiVersion1 = "/api/v1/"

// Router contains the router and mongo instance
type Router interface {
	GET(path string, handle httprouter.Handle)
	POST(path string, handle httprouter.Handle)
	DELETE(path string, handle httprouter.Handle)
	Send(w http.ResponseWriter, response interface{})
	JSON(w http.ResponseWriter, response interface{})
	GetRouter() *httprouter.Router
}

type hayumRouter struct {
	router   *httprouter.Router
	basePath string
}

// NewHayumRouter creates a new Hayum router
func NewHayumRouter(basePath string) Router {
	return &hayumRouter{httprouter.New(), basePath}
}

func (hr *hayumRouter) GET(path string, handle httprouter.Handle) {
	hr.router.GET(util.ConstructEndpoint(hr.basePath, path), handle)
}

func (hr *hayumRouter) POST(path string, handle httprouter.Handle) {
	hr.router.POST(util.ConstructEndpoint(hr.basePath, path), handle)
}

func (hr *hayumRouter) DELETE(path string, handle httprouter.Handle) {
	hr.router.DELETE(util.ConstructEndpoint(hr.basePath, path), handle)
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
func NewRouter() Router {
	router := NewHayumRouter(util.ConstructEndpoint(apiVersion1, "/"))

	initUserRoute(router)
	initS3Route(router)

	return router
}
