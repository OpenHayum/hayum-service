package route

import (
	"fmt"
	"net/http"

	"bitbucket.org/hayum/hayum-service/util"
	"github.com/julienschmidt/httprouter"
)

func (r *Router) UserRoute(basePath string) {
	r.GET(util.ConstructEndpoint(basePath, "/:name"), getUser)
}

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
