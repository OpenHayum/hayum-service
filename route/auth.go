package route

import (
	"encoding/json"
	"net/http"

	"github.com/julienschmidt/httprouter"

	"bitbucket.org/hayum/hayum-service/models"
	"bitbucket.org/hayum/hayum-service/service"
)

type authRoute struct {
	router Router
	s      service.AuthServicer
}

func initAuthRoute(router Router) {

	a := &authRoute{router, service.NewAuthService()}

	a.router.POST("/account", a.createNewAccount)
}

func (a *authRoute) createNewAccount(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	schemaDecoder.SetAliasTag("json")
	var user models.User

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err, response := a.s.Register(&user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	a.router.JSON(w, response)
}
