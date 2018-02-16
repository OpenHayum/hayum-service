package route

import "github.com/julienschmidt/httprouter"

type Router struct {
	*httprouter.Router
}

func (r *Router) Init() {
	r.UserRoute("/user")
}
