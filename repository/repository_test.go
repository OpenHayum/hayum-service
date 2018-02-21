package repository

import (
	"reflect"
	"testing"

	"bitbucket.org/hayum/hayum-service/config"
)

func TestNewRepository(t *testing.T) {
	type args struct {
		mongo          *config.Mongo
		collectionName string
	}
	tests := []struct {
		name string
		args args
		want *Repository
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRepository(tt.args.mongo, tt.args.collectionName); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}
