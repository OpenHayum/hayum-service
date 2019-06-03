package errors

import "errors"

var (
	ErrSessionAlreadyDeleted   = errors.New("session already deleted")
	ErrUnableToGetSession      = errors.New("unable to create session")
	ErrCookieMissingFromHeader = errors.New("cookie not present in header")
)
