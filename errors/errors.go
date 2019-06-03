package errors

import (
	"errors"
	"hayum/core_apis/logger"
	"net/http"
)

var (
	ErrSessionAlreadyDeleted   = errors.New("session already deleted")
	ErrUnableToGetSession      = errors.New("unable to create session")
	ErrCookieMissingFromHeader = errors.New("cookie not present in header")
)

func CheckAndSendResponseErrorWithStatus(err error, w http.ResponseWriter, statusCode int) {
	if err != nil {
		logger.Log.Error(err)
		http.Error(w, err.Error(), statusCode)
		return
	}
}

func CheckAndSendResponseInternalServerError(err error, w http.ResponseWriter) {
	CheckAndSendResponseErrorWithStatus(err, w, http.StatusInternalServerError)
}
