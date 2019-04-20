package main

import (
	"context"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"hayum/core_apis/schema"
	"log"
	"time"
)

func main() {
	ctx := context.Background()

	// Set timeout for 2 sec to connect to the database
	ctx, cancel := context.WithTimeout(ctx, time.Second*2)
	defer cancel()

	db, err := sqlx.ConnectContext(ctx, "mysql", "root:devmysql@/hayum?multiStatements=true")
	if err != nil {
		log.Fatalln(err)
	}

	ctx, cancel = context.WithCancel(ctx)
	time.AfterFunc(time.Second, cancel)

	db.MustExecContext(ctx, schema.DDL)
}
