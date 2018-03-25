package main

import (
	"io"
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
	log.Println("GO_ENV:", env)

	internalConfigDetail := config.NewDetail("./config", "config")
	externalConfigDetail := config.NewDetail(config.ExternalConfigFilePath, "external_config")

	var err error
	_, err = config.LoadConfig(internalConfigDetail, externalConfigDetail)

	if err != nil {
		log.Panic(err)
	}

	port = config.App.GetString("port")

	switch env {
	case devEnv:
		dbURL = config.App.GetString("db.dev.url")
		dbName = config.App.GetString("db.dev.name")
	case prodEnv:
		dbURL = config.App.GetString("db.prod.url")
		dbName = config.App.GetString("db.prod.name")
	default:
		dbURL = "localhost"
		dbName = "hayum"
	}
}

func initLogger() {
	logpath := config.App.GetString("logpath")
	f, err := os.OpenFile(logpath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Panic(err)
	}
	defer f.Close()
	mw := io.MultiWriter(os.Stdout, f)
	log.SetOutput(mw)
	log.Printf("Using log path: %s\n", logpath)
}

func main() {
	initConfig()
	initLogger()

	router := route.NewRouter(dbURL, dbName)

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router.GetRouter())

	log.Println("Listening on port:", port)

	log.Panic(http.ListenAndServe(port, middleware))
}
