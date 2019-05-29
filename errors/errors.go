package errors

import "errors"

var (
	ErrSessionAlreadyDeleted = errors.New("Session already deleted")
)
