package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func Connection() *sql.DB {
	db, _ := sql.Open("mysql", "root:dindarani1@tcp(localhost:3306)/artikel")

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db
}
