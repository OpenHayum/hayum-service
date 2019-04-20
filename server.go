package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"hayum/core_apis/schema"
	"log"
)

func main() {
	db, err := sqlx.Connect("mysql", "root:devmysql@/hayum?multiStatements=true")
	if err != nil {
		log.Fatalln(err)
	}

	//ctx := context.Background()

	db.MustExec(schema.DDL)
}
