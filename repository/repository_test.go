package repository

import (
	"errors"
	"reflect"
	"testing"

	"github.com/stretchr/testify/mock"
	"gopkg.in/mgo.v2"
)

type repositoryMock struct {
	mock.Mock
}

func (m *repositoryMock) Save(model interface{}) error {
	args := m.Called(model)
	return args.Error(0)
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
	type fields struct {
		collection *mgo.Collection
	}
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    interface{}
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &Repository{
				collection: tt.fields.collection,
			}
			got, err := r.GetByID(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Repository.GetByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Repository.GetByID() = %v, want %v", got, tt.want)
			}
		})
	}
}
