package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/urfave/negroni"
)

const (
	devEnv  = "dev"
	prodEnv = "prod"
	testEnv = "test"
)

var dbName, dbURL, port string

func initConfig() {
	env := os.Getenv("GO_ENV")
	fmt.Println("GO_ENV:", env)

	internalConfigDetail := config.NewDetail("./config", "config")
	externalConfigDetail := config.NewDetail(config.ExternalConfigFilePath, "external_config")

	appConfig, err := config.LoadConfig(internalConfigDetail, externalConfigDetail)

	if err != nil {
		log.Panic(err)
	}

	port = appConfig.GetString("port")

	switch env {
	case devEnv:
		dbURL = appConfig.GetString("db.dev.url")
		dbName = appConfig.GetString("db.dev.name")
	case prodEnv:
		dbURL = appConfig.GetString("db.prod.url")
		dbName = appConfig.GetString("db.prod.name")
	default:
		dbURL = "localhost"
		dbName = "hayum"
	}
}

func main() {
	initConfig()

	router := route.NewRouter(dbURL, dbName)

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router.GetRouter())

	fmt.Println("Listening", port)

	log.Panic(http.ListenAndServe(port, middleware))
}
