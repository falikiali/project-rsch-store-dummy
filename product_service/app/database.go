package app

import (
	"database/sql"
	"fmt"
	"rsch/product_service/app/config"
	"rsch/product_service/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(config *config.Config) *sql.DB {
	dbSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", config.Database.User, config.Database.Password, config.Database.Host, config.Database.Port, config.Database.Name)
	db, err := sql.Open("mysql", dbSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxIdleTime(10 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	return db
}
