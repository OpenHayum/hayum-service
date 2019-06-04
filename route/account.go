package route

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"hayum/core_apis/errors"
	"hayum/core_apis/middleware"
	"hayum/core_apis/models"
	"hayum/core_apis/repository"
	"hayum/core_apis/service"
	"net/http"
	"strconv"
)

type accountRoute struct {
	router  Router
	service service.AccountService
}

// Account routes will be protected [requires logged in users]
func initAccountRoute(router Router) {
	accountRepo := repository.NewSQLAccountRepository(router.GetConn())
	accountService := service.NewAccountService(accountRepo)

	accRoute := &accountRoute{router, accountService}

	accRoute.router.POST("/account", middleware.Protected(accRoute.createAccount))
	accRoute.router.GET("/account/:id", middleware.Protected(accRoute.getAccount))
}

func (ar *accountRoute) createAccount(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	acct := new(models.Account)
	err := json.NewDecoder(r.Body).Decode(acct)
	if errors.CheckAndSendResponseBadRequest(err, w) {
		return
	}

	err = ar.service.Save(r.Context(), acct)
	if errors.CheckAndSendResponseInternalServerError(err, w) {
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (ar *accountRoute) getAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	idStr := ps.ByName("id")
	acctId, err := strconv.Atoi(idStr)
	if errors.CheckAndSendResponseBadRequest(err, w) {
		return
	}
	acct, err := ar.service.GetByID(r.Context(), acctId)

	ar.router.JSON(w, acct)
}
