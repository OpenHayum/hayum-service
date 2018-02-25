package main

import (
	"fmt"
	"log"
	"net/http"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/urfave/negroni"
)

func main() {
	appConfig, err := config.LoadConfig("./config")
	if err != nil {
		log.Panic(err)
	}

	port := appConfig.GetString("port")
	dbURL := appConfig.GetString("dev_db_url")
	dbName := appConfig.GetString("dev_db_name")
	fmt.Println("Listening", port)

	router := route.NewRouter(dbURL, dbName)

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router)

	http.ListenAndServe(port, middleware)
}
