package config

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB_USER     string
	DB_PASSWORD string
	DB_HOST     string
)

func init() {
	DB_USER = os.Getenv("DBUSER")
	DB_PASSWORD = os.Getenv("DBPASSWORD")
	DB_HOST = os.Getenv("DBHOST")
}

func NewDB() *sql.DB {
	dbConfig := fmt.Sprintf("%s:%s@tcp(%s)/%s", DB_USER, DB_PASSWORD, DB_HOST, "sfs")

	dbtx, err := sql.Open("mysql", dbConfig)
	if err != nil {
		log.Fatalln(err)
	}
	return dbtx
}
