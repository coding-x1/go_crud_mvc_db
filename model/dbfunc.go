package model

import (
	"database/sql"
	"log"
)

func Create(val string) {
	db := connect()
	defer db.Close()
	_, err := db.Query("INSERT into entryval(entryval) VALUES ($1)", val)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
}
func Delete(id1 int) {
	db := connect()
	defer db.Close()
	_, err := db.Query("Delete from entryval where id =$1", id1)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
}
func Update(val string, id1 int) {
	db := connect()
	defer db.Close()
	_, err := db.Query("Update entryval set entryval=$1 where id = $2", val, id1)
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
}
func Index() *sql.Rows {
	db := connect()
	defer db.Close()
	rows, err := db.Query("Select * from entryval")
	if err != nil {
		log.Fatalf("An error occured while executing query: %v", err)
	}
	return rows
}
