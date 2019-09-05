package infrastructure

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func init() {
	var err error
	DB, err = sql.Open("mysql", "user:password@tcp(0.0.0.0)/sampleDB?parseTime=true")

	if err != nil {
		log.Fatal(err)
	}
}
