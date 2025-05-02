package config

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connstr := "postgres://postgres:admin@localhost:5433/login-service?sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connstr)
	if err != nil {
		log.Fatal("connection error", err)
	}

	if err := DB.Ping(); err != nil {
		log.Fatal("Database unreachable", err)
	}

	log.Println("Connection established")

}
