package db

import (
	"context"
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
	"log"
	"time"
)

type Conn struct {
	*sqlx.DB
}

func getDSN(cfg *viper.Viper) string {
	user := cfg.GetString("db.user")
	password := cfg.GetString("db.password")
	db := cfg.GetString("db.name")

	return fmt.Sprintf("%s:%s@/%s?multiStatements=true&parseTime=true", user, password, db)
}

func OpenContext(ctx context.Context, cfg *viper.Viper) *Conn {
	// connect to database
	sqlx.NameMapper = func(s string) string { return s }
	db, err := sqlx.ConnectContext(ctx, "mysql", getDSN(cfg))
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	// create mysql tables
	db.MustExecContext(ctx, createDDL)

	return &Conn{db}
}
