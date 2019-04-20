package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/rs/cors"
	"github.com/urfave/negroni"
	"hayum/core_apis/config"
	"log"
	"net/http"
	"time"
)

func main() {
	cfg := config.New()

	ctx := context.Background()

	// Set timeout for 2 sec to connect to the database
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	dbOpenContext(ctx, cfg)

	// setup middleware
	middleware := negroni.New()
	middleware.Use(negroni.NewLogger())

	handler := cors.Default().Handler(middleware)
	port := cfg.GetString("port")

	log.Println("Listening on port:", port)
	log.Panic(http.ListenAndServe(port, handler))

}
