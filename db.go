package main

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"hayum/core_apis/schema"
	"log"
	"time"
)

func getDSN(cfg *viper.Viper) string {
	user := cfg.GetString("db.user")
	password := cfg.GetString("db.password")
	db := cfg.GetString("db.name")

	return fmt.Sprintf("%s:%s@/%s?multiStatements=true", user, password, db)
}

func dbOpenContext(ctx context.Context, cfg *viper.Viper) *sqlx.DB {
	// connect to database
	db, err := sqlx.ConnectContext(ctx, "mysql", getDSN(cfg))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	// create mysql tables
	db.MustExecContext(ctx, schema.DDL)

	return db
}
