package configs

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func ConnectDatabase() *sql.DB {
	connection := "user=postgres dbname=go_crud password=postgres host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {
		log.Fatalln("Error on connect db", err)
	}

	return db
}
