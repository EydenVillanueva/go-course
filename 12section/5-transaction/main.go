package main

// Database transaction
// 1. User creates account
// 2. Create a wallet for the user
// 3. Want to top up the wallet for the user
// 4. You want to write a transaction log
import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/mattn/go-sqlite3" // (blank import) we are not going to use anything from here
	"golang.org/x/crypto/bcrypt"
)

var schema = `
	CREATE TABLE IF NOT EXISTS profile (
		user_id INTEGER PRIMARY KEY REFERENCES users(user_id) ON DELETE CASCADE,
		avatar TEXT NOT NULL,
		created_at DATETIME NOT NULL DEFAULT CURRENT_TIMESTAMP
	)
`

type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
	Profile        Profile   `json:"profile"`
}

type Profile struct {
	id      int
	Avatar  string
	Created time.Time
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

	userID, err := createUser(db, "user from tx", "tx@localhost", "alsdkfaf", "http://avatar/user")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(userID)

}

func createTable(db *sql.DB) {
	_, err := db.Exec(schema)
	if err != nil {
		log.Fatal(err)
	}
}

// Begin, RollBack or Commit
func createUser(db *sql.DB, name, email, hashed_password, avatar string) (int64, error) {
	ctx := context.Background()
	tx, err := db.BeginTx(ctx, nil)

	if err != nil {
		log.Fatal(err)
	}

	stmt, err := tx.PrepareContext(ctx, `INSERT INTO users (name, email, hashed_password) VALUES (?,?,?)`)

	if err != nil {
		return 0, err
	}

	defer stmt.Close()

	hp, err := bcrypt.GenerateFromPassword([]byte(hashed_password), bcrypt.DefaultCost)

	if err != nil {
		return 0, err
	}

	result, err := stmt.Exec(name, email, string(hp))

	if err != nil {
		log.Fatal(err)
	}

	userId, err := result.LastInsertId()
	if err != nil {
		return 0, err
	}

	profileStm, err := tx.PrepareContext(ctx, `INSERT INTO profile (user_id, avatar) VALUES(?,?)`)

	if err != nil {
		txError := tx.Rollback()

		if txError != nil {
			return 0, txError
		}
		return 0, err
	}

	err = tx.Commit()

	if err != nil {
		return 0, err
	}

	defer profileStm.Close()

	result, err = profileStm.Exec(userId, avatar)
	return userId, nil
}

func GetUserByEmail(db *sql.DB, email string) (*User, error) {
	var stmt = `SELECT u.id, u.name, u.email, u.hashed_password, u.created_at FROM users u 
	INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ? LIMIT 1`

	row := db.QueryRow(stmt, email)

	// Scanning
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}
