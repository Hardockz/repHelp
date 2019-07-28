package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	host     = "172.19.0.2"
	port     = 5432
	user     = "postgres"
	password = "password"
	dbname   = "postgres"
)

var schema = `
	CREATE TABLE projects (
		id text PRIMARY KEY,
		name text NOT NULL UNIQUE,
		status text NOT NULL 
);`

func createConnString() string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	return psqlInfo
}

func createDB() *sqlx.DB {
	db, err := sqlx.Open("postgres", createConnString())
	if err != nil {
		log.Fatal(err)
	}

	db.Exec(schema)

	//defer db.Close()

	return db
}
