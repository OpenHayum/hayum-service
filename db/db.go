package db

import (
	"context"
	"fmt"
	"hayum/core_apis/logger"
	"log"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Conn struct {
	*sqlx.DB
}

func getDSN(cfg *viper.Viper) string {
	user := cfg.GetString("db.user")
	password := cfg.GetString("db.password")
	db := cfg.GetString("db.name")
	address := cfg.GetString("db.url")

	logger.Log.Info("Address: ", address)
	// config := mysql.Config{
	// 	User:   user,
	// 	Passwd: password,
	// 	Addr:   address,
	// 	DBName: db,
	// }

	// return config.FormatDSN()
	return fmt.Sprintf("%s:%s@%s/%s?multiStatements=true&parseTime=true", user, password, address, db)
}

func OpenContext(ctx context.Context, cfg *viper.Viper) *sqlx.DB {
	// connect to database
	sqlx.NameMapper = func(s string) string { return s }
	dsn := getDSN(cfg)
	logger.Log.Info("DSN: ", dsn)
	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	// create mysql tables
	db.MustExecContext(ctx, createDDL)

	return db
}
