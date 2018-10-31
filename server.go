package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"bitbucket.org/hayum/hayum-service/config"
	"bitbucket.org/hayum/hayum-service/db"
	hayumMiddleware "bitbucket.org/hayum/hayum-service/middleware"
	"bitbucket.org/hayum/hayum-service/route"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

const (
	devEnv     = "development"
	prodEnv    = "production"
	stagingEnv = "staging"
	testEnv    = "test"
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

	config.App.Set("dbName", dbName)
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
	hayumMiddleware.InitMiddlewareServices()

	err := db.NewMongoSession(dbURL, dbName)
	if err != nil {
		log.Println("Unable to connect to mongo")
		log.Panic(err.Error())
	}

	router := route.NewRouter()

	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())
	middleware.UseHandler(router.GetRouter())
	middleware.UseHandlerFunc(hayumMiddleware.Authorize)

	log.Println("Listening on port:", port)

	handler := cors.Default().Handler(middleware)

	log.Panic(http.ListenAndServe(port, handler))
}
