package route

import (
	"encoding/json"
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/models"
	"github.com/gorilla/schema"

	"bitbucket.org/hayum/hayum-service/service"
	"github.com/julienschmidt/httprouter"
)

type accountRoute struct {
	router  Router
	service service.AccountServicer
}

var schemaDecoder = schema.NewDecoder()

func initAccountRoute(router Router) {
	s := service.NewAccountService()
	a := &accountRoute{router, s}

	a.router.POST("/account", a.createNewAccount)
	a.router.GET("/account/:id", a.getAccountByID)
	a.router.POST("/account/:id", a.updateAccount)
}

func (ar *accountRoute) createNewAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var account models.Account
	userID := r.Header.Get("user-id")

	if userID == "" {
		log.Println("createNewAccount: user-id not present in header")
		http.Error(w, "UserID is required in header", http.StatusBadRequest)
		return
	}

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		log.Println("createNewAccount: unable to decode account request body", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ar.service.CreateNewAccount(userID, &account); err != nil {
		log.Println("createNewAccount: unable to create new account from request body", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	ar.router.JSON(w, account)
}

func (ar *accountRoute) getAccountByID(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	acc := &models.Account{}
	err := ar.service.GetAccountByID(ps.ByName("id"), acc)

	if err != nil {
		log.Println("Cannot get account by id", err.Error())
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	ar.router.JSON(w, acc)
}

func (ar *accountRoute) updateAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		log.Println("createNewAccount: unable to decode account request body", err.Error())
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if err := ar.service.UpdateAccount(&account); err != nil {
		log.Println("createNewAccount: unable to update account from request body", err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
	ar.router.Send(w, "Updated")
}
