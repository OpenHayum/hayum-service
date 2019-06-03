package db

import (
	"context"
	"fmt"
	"github.com/srinathgs/mysqlstore"
	"hayum/core_apis/logger"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql" // SQL connect context requires mysql dialect
	"github.com/jmoiron/sqlx"
	"github.com/spf13/viper"
)

type Conn struct {
	*sqlx.DB
}

var Store *mysqlstore.MySQLStore

func getDSN(cfg *viper.Viper) string {
	user := cfg.GetString("db.user")
	password := cfg.GetString("db.password")
	db := cfg.GetString("db.name")
	address := cfg.GetString("db.url")

	return fmt.Sprintf("%s:%s@%s/%s?multiStatements=true&parseTime=true", user, password, address, db)
}

func createSessionStore(db *sqlx.DB, cfg *viper.Viper) {
	var err error

	tableName := cfg.GetString("session.table_name")
	maxAge := cfg.GetInt("session.expiration_time_ms")
	secretKey := cfg.GetString("session.secret_key")

	Store, err = mysqlstore.NewMySQLStoreFromConnection(db.DB, tableName, "/", maxAge, []byte(secretKey))
	if err != nil {
		logger.Log.Panic(err)
	}
}

func OpenContext(ctx context.Context, cfg *viper.Viper) *sqlx.DB {
	sqlx.NameMapper = func(s string) string { return s }
	dsn := getDSN(cfg)
	logger.Log.Info("DSN: ", dsn)
	db, err := sqlx.ConnectContext(ctx, "mysql", dsn)
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	// create mysql tables
	db.MustExecContext(ctx, createDDL)
	createSessionStore(db, cfg)

	return db
}
