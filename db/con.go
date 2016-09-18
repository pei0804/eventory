package db

import (
	"database/sql"
	"log"
)

func ConDB() (db *sql.DB) {
	db, err := sql.Open("mysql", "root:password@/eventory")
	if err != nil {
		log.Fatal("open erro: %v", err)
	}
	return db
}
