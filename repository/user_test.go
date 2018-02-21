package repository

import (
	"reflect"
	"testing"

	"bitbucket.org/hayum/hayum-service/models"
)

func TestNewUserRepository(t *testing.T) {
	type args struct {
		r *Repository
	}
	tests := []struct {
		name string
		args args
		want *UserRepository
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUserRepository(tt.args.r); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUserRepository() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_CreateNewUser(t *testing.T) {
	type args struct {
		user *models.User
	}
	tests := []struct {
		name    string
		r       *UserRepository
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.CreateNewUser(tt.args.user); (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.CreateNewUser() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUserRepository_GetUserByID(t *testing.T) {
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		r       *UserRepository
		args    args
		want    *models.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUserByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserRepository_GetUserByEmail(t *testing.T) {
	type args struct {
		email string
	}
	tests := []struct {
		name    string
		r       *UserRepository
		args    args
		want    *models.User
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetUserByEmail(tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("UserRepository.GetUserByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserRepository.GetUserByEmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
