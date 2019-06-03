package test

import (
	"hayum/core_apis/db"
	"hayum/core_apis/route"
	"net/http/httptest"
)

func newServer(conn *db.Conn) *httptest.Server {
	r := route.NewRouter(conn)
	return httptest.NewServer(r.GetRouter())
}
