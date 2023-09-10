package config

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	dsn = "%s:%s@tcp(%s:%v)/%s?charset=utf8mb4&parseTime=True&loc=Local"
)

func MysqlConnection(config Config) *sql.DB {
	db, err := sql.Open(config.Get("db_driver"), fmt.Sprintf(dsn, config.Get("db_username"), config.Get("db_password"), config.Get("db_url"), config.Get("db_port"), config.Get("db_name")))
	if err != nil {
		log.Fatalf(err.Error())
	}

	return db
}
