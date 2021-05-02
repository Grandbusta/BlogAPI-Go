package db

import (
	"database/sql"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func DbConfig() (*sql.DB, error) {
	DB_NAME := os.Getenv("DB_NAME")
	DB_PASS := os.Getenv("DB_PASS")
	DB_USER := os.Getenv("DB_USER")
	DB_HOST := os.Getenv("DB_HOST")
	DB_PORT := os.Getenv("DB_PORT")
	db, err := sql.Open("mysql", DB_USER+":"+DB_PASS+"@tcp("+DB_HOST+":"+DB_PORT+")/"+DB_NAME)
	if err != nil {
		return nil, err
	}
	return db, nil
}
