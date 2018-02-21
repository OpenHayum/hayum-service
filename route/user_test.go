package route

import (
	"net/http"
	"testing"

	"bitbucket.org/hayum/hayum-service/service"
	"github.com/julienschmidt/httprouter"
)

func Test_initUserRoute(t *testing.T) {
	type args struct {
		router   *Router
		basePath string
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			initUserRoute(tt.args.router, tt.args.basePath)
		})
	}
}

func Test_userRoute_createUser(t *testing.T) {
	type fields struct {
		router  *Router
		service *service.UserService
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRoute{
				router:  tt.fields.router,
				service: tt.fields.service,
			}
			u.createUser(tt.args.w, tt.args.r, tt.args.ps)
		})
	}
}

func Test_userRoute_getUser(t *testing.T) {
	type fields struct {
		router  *Router
		service *service.UserService
	}
	type args struct {
		w  http.ResponseWriter
		r  *http.Request
		ps httprouter.Params
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := &userRoute{
				router:  tt.fields.router,
				service: tt.fields.service,
			}
			u.getUser(tt.args.w, tt.args.r, tt.args.ps)
		})
	}
}
