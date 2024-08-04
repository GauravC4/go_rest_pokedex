package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/jackc/pgx/v5/stdlib"
)

func createDbConn() *sql.DB {
	db, err := sql.Open("pgx", "postgres://postgres:postgrespw@localhost:32768/pokedexdb")
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	return db
}
