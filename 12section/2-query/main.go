package main

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // (blank import) we are not going to use anything from here
	"golang.org/x/crypto/bcrypt"
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

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

func main() {
	dbName := "users_database.db"
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

	// createTable(db)
	// userNames := []string{"Javier", "Paula", "Guillermo", "Carolina", "Helen"}

	// for _, username := range userNames {
	// 	email := fmt.Sprintf(`%s@gmail.com`, username)
	// 	lastID, err := createUser(db, username, email, "password")
	// 	fmt.Println("last user is is", lastID)
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// }

	user, err := GetUserByEmail(db, "Paula@gmail.com")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("paula is %v\n", user)

	fmt.Println("Connected...")
}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

func createUser(db *sql.DB, name, email, hashed_password string) (int64, error) {
	stmt := `INSERT INTO users (name, email, hashed_password) VALUES (?,?,?)`

	hp, err := bcrypt.GenerateFromPassword([]byte(hashed_password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	result, err := db.Exec(stmt, name, email, string(hp))

	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var stmt = `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)

	// Scanning
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
