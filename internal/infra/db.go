package infra

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func CreateDBConnection() *sql.DB {
	db, err := sql.Open("postgres", "user=docker dbname=apidev sslmode=disable password=docker host=localhost")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	if err := db.Ping(); err != nil {
		panic(err)
	} else {
		fmt.Println("Successfuly connected")
	}

	return db
}
