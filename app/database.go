package app

import (
	"database/sql"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func NewDb() *sql.DB {
	errEnv := godotenv.Load(".env")

	if errEnv != nil {
		panic(errEnv)
	}

	username := os.Getenv("username")
	password := os.Getenv("password")
	db_name := os.Getenv("db_name")
	db, err := sql.Open("mysql", username+":"+password+"@tcp(localhost:3306)/"+db_name)

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(10)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(10 * time.Minute)

	if err != nil {
		panic(err)
	}

	return db

}
