package main

import (
	"context"
	"hayum/core_apis/config"
	"hayum/core_apis/db"
	"hayum/core_apis/logger"
	route "hayum/core_apis/routes"
	"net/http"
	"time"

	"github.com/rs/cors"
	"github.com/urfave/negroni"
)

func main() {
	logger.Init()
	cfg := config.New()
	ctx := context.Background()
	// Set timeout for 2 sec to connect to the database
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()
	conn := db.OpenContext(ctx, cfg)
	router := route.NewRouter(&db.Conn{DB: conn})

	// setup middleware
	middleware := negroni.New()
	middleware.UseHandler(router.GetRouter())
	middleware.Use(negroni.NewLogger())

	handler := cors.Default().Handler(middleware)
	port := cfg.GetString("port")
	logger.Log.Info("Listening on port:", port)

	logger.Log.Panic(http.ListenAndServe(port, handler))
}
