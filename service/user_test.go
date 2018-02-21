package service

import (
	"reflect"
	"testing"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/models"
)

func TestNewUserService(t *testing.T) {
	type args struct {
		mongo *config.Mongo
	}
	tests := []struct {
		name string
		args args
		want *UserService
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserService(tt.args.mongo); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_CreateNewUser(t *testing.T) {
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		s       *UserService
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.s.CreateNewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserService.CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserService_GetUserByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		s       *UserService
		args    args
		want    *models.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserService_GetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		s       *UserService
		args    args
		want    *models.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.s.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserService.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserService.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
