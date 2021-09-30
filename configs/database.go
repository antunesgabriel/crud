package configs

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDatabase() (*sql.DB, error) {
	connection := "user=postgres dbname=go_crud password=postgres host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgres", connection)

	if err != nil {

		return nil, err
	}

	return db, nil
}
