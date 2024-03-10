package app

import (
	"database/sql"
	"os"
	"rsch/auth_service/helper"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDB() *sql.DB {
	err := godotenv.Load()
	helper.PanicIfError(err)

	db, err := sql.Open(os.Getenv("DB"), os.Getenv("DB_USER")+":"+os.Getenv("DB_PASSWORD")+"@"+os.Getenv("DB_PROTOCOL")+"("+os.Getenv("DB_ADDRESS")+")/"+os.Getenv("DB_NAME"))
	helper.PanicIfError(err)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(20)
	db.SetConnMaxLifetime(60 * time.Minute)
	db.SetConnMaxIdleTime(10 * time.Minute)

	return db
}
