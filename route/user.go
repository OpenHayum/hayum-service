package route

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}
