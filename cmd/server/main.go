package main

import (
	"database/sql"
	"fmt"
	"log"

	"net/http"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi, I love %s", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", rootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Setup the database,
func run() error {
	connectionString := "data/goblog_test.db"

	//setup database connection
	db, err = setupDatabase(connectionString)

	if err != nil {
		return nil
	}

	return nil
}

// Setup the database connection
func setupDatabase(connString string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", connString)
	if err != nil {
		return nil, err
	}

	err = db.Ping()

	if err != nil {
		return nil, err
	}

	return db, nil
}
