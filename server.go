package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	fmt.Fprint(w, "Welcome!\n")
}

func Hello(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	fmt.Fprintf(w, "hello, %s!\n", ps.ByName("name"))
}

func main() {
	fmt.Println("Listening", config.Port)

	r := route.Router{httprouter.New()}
	r.Init()

	n := negroni.New()
	n.Use(negroni.NewLogger())
	n.UseHandler(r)

	http.ListenAndServe(config.Port, n)
}
