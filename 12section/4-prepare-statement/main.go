package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // (blank import) we are not going to use anything from here
	"golang.org/x/crypto/bcrypt"
)

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

	// users, err := GetUsers(db)
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// bs, err := json.MarshalIndent(users, "", " ")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(bs))

	ctx := context.Background()

	_, err = createUserWithPrepared(ctx, db, "Francisco", "francisco@gmail.com", "test091283")

	if err != nil {
		log.Fatal(err)
	}
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var stmt = `SELECT id, name, email, hashed_password, created_at FROM users WHERE email = ?`

	row := db.QueryRow(stmt, email)
	var user User

	// Scanning
	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetUsers(db *sql.DB) ([]User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`
	rows, err := db.Query(stmt)

	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	users := []User{}

	//Scanning
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID,
			&user.Name,
			&user.Email,
			&user.HashedPassword,
			&user.CreatedAt,
		); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
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

func createUserWithPrepared(ctx context.Context, db *sql.DB, name, email, hashed_password string) (int64, error) {
	stmt, err := db.Prepare(`INSERT INTO users (name, email, hashed_password) VALUES (?,?,?)`)

	if err != nil {
		return 0, err
	}

	hp, err := bcrypt.GenerateFromPassword([]byte(hashed_password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	result, err := stmt.ExecContext(ctx, name, email, string(hp))

	if err != nil {
		log.Fatal(err)
	}

	return result.LastInsertId()
}
