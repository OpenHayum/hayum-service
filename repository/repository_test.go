package repository

import (
	"errors"
	"reflect"
	"testing"

	"bitbucket.org/hayum/hayum-service/models"

	"github.com/stretchr/testify/mock"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(model interface{}) error {
	args := m.Called(model)
	return args.Error(0)
}

func (m *repositoryMock) GetByID(id string) (interface{}, error) {
	args := m.Called(id)
	return args.Get(0), args.Error(1)
}

func TestRepository_Save(t *testing.T) {
	m := &struct{ name string }{name: "asem"}
	mock := new(repositoryMock)
	tests := []struct {
		name    string
		err     error
		wantErr bool
	}{
		{
			name:    "Test Repository Save fail",
			err:     errors.New("Unable to save"),
			wantErr: true,
		},
		{
			name:    "Test Repository Save success",
			err:     nil,
			wantErr: false,
		},
	}

	for _, tt := range tests {
		mock.On("Save", m).Return(tt.err)

		t.Run(tt.name, func(t *testing.T) {
			t.Log(tt.err, tt.wantErr)

			if err := mock.Save(m); (err == tt.err) != tt.wantErr {
				t.Errorf("Repository.Save() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}

	mock.AssertExpectations(t)
}

func TestRepository_GetByID(t *testing.T) {
	mock := new(repositoryMock)
	tests := []struct {
		name    string
		id      string
		err     error
		wantErr bool
	}{
		{
			name:    "Test Repository GetByID fail",
			id:      "id",
			err:     errors.New("Unable to save"),
			wantErr: true,
		},
		{
			name:    "Test Repository GetByID success",
			id:      "id",
			err:     nil,
			wantErr: false,
		},
	}

	u := new(models.User)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mock.On("GetByID", tt.id).Return(u, tt.err)
			got, err := mock.GetByID(tt.id)
			if (err == tt.err) != tt.wantErr {
				t.Errorf("Repository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, u) {
				t.Errorf("Repository.GetByID() = %v, want %v", got, u)
			}
		})
	}
}
