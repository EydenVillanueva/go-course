package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/mattn/go-sqlite3" // (blank import) we are not going to use anything from here
)

var schema = `
	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		email TEXT NOT NULL UNIQUE,
		hashed_password BLOB NOT NULL, --storing as BLOB for byte slice
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
`

func main() {
	dbName := "data.db"
	_ = os.Remove(dbName)

	db, err := sql.Open("sqlite3", dbName)

	if err != nil {
		log.Fatal(err)
	}

	defer func() {
		fmt.Println("Closing database")
		if err := db.Close(); err != nil {
			log.Printf("error closing database connection: %v", err)
		}
	}()

	err = db.Ping()

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected...")

	_, err = db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("table was created")

}
