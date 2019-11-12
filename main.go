package main

import (
	"database/sql"
	"log"
	"net/http"
)

var db *sql.DB

func init() {
	tmpDB, err := sql.Open("postgres", "dbname=UsersDB user=postgres password=localhost host=localhost sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	db = tmpDB
}

func main() {

	// http.HandleFunc("/", handleSaveUsers)
	http.HandleFunc("/save", handleSaveUsers)
	log.Fatal(http.ListenAndServe(":8090", nil))
}
