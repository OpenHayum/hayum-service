package main

import (
	"fmt"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	fmt.Println("Listening", config.Port)

	router := route.Router{httprouter.New()}
	router.Init()

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router)

	http.ListenAndServe(config.Port, middleware)
}
