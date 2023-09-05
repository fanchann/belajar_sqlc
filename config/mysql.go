package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = "%s:%s@/%s"
)

func MysqlConnection(config Config) *sql.DB {
	db, err := sql.Open(config.Get("db_driver"), fmt.Sprintf(dsn, config.Get("db_username"), config.Get("db_password"), config.Get("db_name")))
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}
