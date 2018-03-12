package db

import (
	"reflect"
	"testing"
)

func TestNewMongoSession(t *testing.T) {
	type args struct {
		url string
		db  string
	}
	tests := []struct {
		name    string
		args    args
		want    *Mongo
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMongoSession(tt.args.url, tt.args.db)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMongoSession() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewMongoSession() = %v, want %v", got, tt.want)
			}
		})
	}
}
