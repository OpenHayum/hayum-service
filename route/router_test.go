package route

import (
	"net/http"
	"testing"

	"bitbucket.org/hayum/hayum-service/config"
	"github.com/julienschmidt/httprouter"
)

func TestRouter_Init(t *testing.T) {
	type fields struct {
		Router *httprouter.Router
		Mongo  *config.Mongo
	}
	tests := []struct {
		name   string
		fields fields
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Router{
				Router: tt.fields.Router,
				Mongo:  tt.fields.Mongo,
			}
			r.Init()
		})
	}
}

func TestRouter_Send(t *testing.T) {
	type fields struct {
		Router *httprouter.Router
		Mongo  *config.Mongo
	}
	type args struct {
		w        http.ResponseWriter
		response interface{}
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
			r := &Router{
				Router: tt.fields.Router,
				Mongo:  tt.fields.Mongo,
			}
			r.Send(tt.args.w, tt.args.response)
		})
	}
}
