package route

import (
	"net/http"
	"reflect"
	"testing"

	"bitbucket.org/hayum/hayum-service/config"
	"github.com/julienschmidt/httprouter"
)

func TestNewRouter(t *testing.T) {
	type args struct {
		dbURL  string
		dbName string
	}
	tests := []struct {
		name string
		args args
		want *Router
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRouter(tt.args.dbURL, tt.args.dbName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRouter() = %v, want %v", got, tt.want)
			}
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
