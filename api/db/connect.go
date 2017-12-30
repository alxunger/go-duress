package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Connect establishes a connection with the default database
func Connect() {
	connStr := "user=duress dbname=duress"
	_, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection with database `duress` established")
	}
}
