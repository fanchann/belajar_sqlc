package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	driver = "mysql"
	dsn    = "fanchann:root@/sqlc"
)

func MysqlConnection() *sql.DB {
	db, err := sql.Open(driver, dsn)
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}
