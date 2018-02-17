package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/julienschmidt/httprouter"
	"github.com/urfave/negroni"
)

func main() {
	appConfig, err := config.LoadConfig("./config")
	if err != nil {
		log.Panic(err)
	}

	port := appConfig.GetString("port")
	fmt.Println("Listening", port)

	router := route.Router{httprouter.New()}
	router.Init()

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router)

	http.ListenAndServe(port, middleware)
}
