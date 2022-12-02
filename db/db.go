package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectDB() *sql.DB {
	constr := "user=postgres dbname=golojadb password=postgres host=localhost sslmode=disable"

	db, err := sql.Open("postgres", constr)
	if err != nil {
		panic(err.Error())
	}

	return db
}
