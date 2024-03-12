package app

import (
	"database/sql"
	"fmt"
	"rsch/profile_service/app/config"
	"rsch/profile_service/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func NewDB(configDB *config.Config) *sql.DB {
	dbSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", configDB.Database.User, configDB.Database.Password, configDB.Database.Host, configDB.Database.Port, configDB.Database.Name)
	db, err := sql.Open("mysql", dbSourceName)
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
