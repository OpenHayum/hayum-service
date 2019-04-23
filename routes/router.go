package route

import (
	"encoding/json"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	"hayum/core_apis/util"
	"net/http"

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
	JSONWithStatus(w http.ResponseWriter, status int, response interface{})
	GetRouter() *httprouter.Router
	GetConn() *db.Conn
}

type hayumRouter struct {
	router   *httprouter.Router
	basePath string
	conn     *db.Conn
}

// NewHayumRouter creates a new Hayum router
func NewHayumRouter(basePath string, conn *db.Conn) Router {
	return &hayumRouter{httprouter.New(), basePath, conn}
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

func (hr *hayumRouter) GetConn() *db.Conn {
	return hr.conn
}

func (hr *hayumRouter) Send(w http.ResponseWriter, response interface{}) {
	err := json.NewEncoder(w).Encode(response)
	if err != nil {
		logger.Log.Fatal(err)
	}
}

func (hr *hayumRouter) JSON(w http.ResponseWriter, response interface{}) {
	hr.JSONWithStatus(w, http.StatusOK, response)
}

func (hr *hayumRouter) JSONWithStatus(w http.ResponseWriter, status int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	hr.Send(w, response)
}

// NewRouter initializes all routes of the service
func NewRouter(conn *db.Conn) Router {
	router := NewHayumRouter(util.ConstructEndpoint(apiVersion1, "/"), conn)

	initUserRoute(router)
	initSessionRoute(router)
	initAuthRoute(router)

	return router
}
