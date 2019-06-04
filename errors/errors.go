package errors

import (
	"errors"
	"hayum/core_apis/logger"
	"net/http"
)

var (
	ErrSessionAlreadyDeleted              = errors.New("session already deleted")
	ErrUnableToGetSession                 = errors.New("unable to create session")
	ErrCookieMissingFromHeader            = errors.New("cookie not present in header")
	ErrUserMobileOrEmailAlreadyAssociated = errors.New("mobile or email already associated with another user")
)

func CheckAndSendResponseErrorWithStatus(err error, w http.ResponseWriter, statusCode int) bool {
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), statusCode)
		return true
	}

	return false
}

func CheckAndSendResponseInternalServerError(err error, w http.ResponseWriter) bool {
	return CheckAndSendResponseErrorWithStatus(err, w, http.StatusInternalServerError)
}

func CheckAndSendResponseBadRequest(err error, w http.ResponseWriter) bool {
	return CheckAndSendResponseErrorWithStatus(err, w, http.StatusBadRequest)
}
