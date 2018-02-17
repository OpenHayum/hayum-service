package route

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func getItem(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprint(w, ps.ByName("id"))
}
