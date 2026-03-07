package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func Connect() *sql.DB {

	connStr := "host=localhost user=postgres password=postgres dbname=coupons port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connStr)

	if err != nil {
		log.Fatal("Error connecting to the database", err)
	}

	err = db.Ping()

	if err != nil {
		log.Fatal("Error pinging the database", err)
	}

	return db
}
