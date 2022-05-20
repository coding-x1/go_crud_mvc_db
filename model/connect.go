package model

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

func connect() *sql.DB {
	connStr := "postgresql://postgres:secret@localhost:5432/entry?sslmode=disable"

	// Connect to database
	db, _ := sql.Open("postgres", connStr)
	// Initialize the first connection to the database, to see if everything works correctly.
	// Make sure to check the error.
	err := db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connected to the Database....")
	return db
}
