package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var conn *sql.DB

// InitDB establishes a connection with the default database
func InitDB() {
	var err error
	conn, err = sql.Open("postgres", "user=duress dbname=duress sslmode=disable")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("connection with database `duress` established")
	}

	if err = conn.Ping(); err != nil {
		log.Panic(err)
	}
}

// TotalCodes return the number of total duress codes
func TotalCodes() int64 {
	var ret int64
	var err error

	rows, err := conn.Query(`SELECT COUNT(*) FROM codes;`)
	if err != nil {
		log.Panic(`Error running the specified query`)
	}
	defer rows.Close()

	for rows.Next() {
		rows.Scan(&ret)
	}

	return ret
}
