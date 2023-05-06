package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var (
	DB_USER     string
	DB_PASSWORD string
	PG_SSL_MODE string
	DB_HOST     string
	APP_ENV     string
)

func init() {
	DB_USER = os.Getenv("PGUSER")
	DB_PASSWORD = os.Getenv("PGPASSWORD")
	PG_SSL_MODE = os.Getenv("PGSSLMODE")
	DB_HOST = os.Getenv("PGHOST")
	APP_ENV = os.Getenv("APP_ENV")
}

func NewDB() *sql.DB {
	dbConfig := fmt.Sprintf("user=%s password=%s host=%s dbname=%s sslmode=%s", DB_USER, DB_PASSWORD, DB_HOST, APP_ENV, PG_SSL_MODE)

	dbtx, err := sql.Open("postgres", dbConfig)
	if err != nil {
		log.Fatalln(err)
	}
	return dbtx
}
