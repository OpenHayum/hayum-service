package util

import (
	"errors"
	"reflect"

	"github.com/fatih/structtag"
)

// TAG is hayum specific tag
const TAG = "hy"

// Validation holds the Message
type Validation struct {
	Message string
}

// NewValidation creates a new Validation
func NewValidation() *Validation {
	return &Validation{}
}

// Validate validates a model
func (v *Validation) Validate(model interface{}) error {
	modelType := reflect.TypeOf(reflect.ValueOf(model))
	numfield := modelType.NumField()

	for i := 0; i < numfield; i++ {
		tag := modelType.Field(i).Tag
		tags, _ := structtag.Parse(string(tag))
		if name, _ := tags.Get(TAG); name.GoString() == "required" {
			v.Message = name.GoString() + " is required"
			return errors.New(v.Message)
		}
	}
	return nil
}
