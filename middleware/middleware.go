package middleware

import (
	"github.com/julienschmidt/httprouter"
	"github.com/pkg/errors"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	hyErrors "hayum/core_apis/errors"
	"net/http"
)

func Protected(next httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		session, err := db.Store.Get(r, config.SessionName)
		hyErrors.CheckAndSendResponseBadRequest(err, w)
		userVal := session.Values["user"]
		//logger.Log.Info("Middleware:", session, userVal, r.Header.Get("Cookie"), session.IsNew)
		if userVal == nil {
			http.Error(w, errors.New("user not logged in").Error(), http.StatusUnauthorized)
			return
		}

		next(w, r, ps)
	}
}
