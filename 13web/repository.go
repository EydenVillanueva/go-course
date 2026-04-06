package main

import (
	"context"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

type UserRepository interface {
	CreateUser(name, email, hashed_password, avatar string) (int64, error)
	GetUserByEmail(email string) (*User, error)
	GetUsers() ([]User, error)
}

type SQLUserRepository struct {
	db *sql.DB
}

func NewSQLUserRepository(db *sql.DB) UserRepository {
	return &SQLUserRepository{db: db}
}

func (r *SQLUserRepository) CreateUser(name, email, hashed_password, avatar string) (int64, error) {
	ctx := context.Background()
	tx, err := r.db.BeginTx(ctx, nil)

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

func (r *SQLUserRepository) GetUserByEmail(email string) (*User, error) {
	var stmt = `SELECT u.id, u.name, u.email, u.hashed_password, u.created_at FROM users u 
	INNER JOIN profile p ON u.id = p.user_id WHERE u.email = ? LIMIT 1`

	row := r.db.QueryRow(stmt, email)

	// Scanning
	var user User

	err := row.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *SQLUserRepository) GetUsers() ([]User, error) {
	stmt := `SELECT id, name, email, hashed_password, created_at FROM users`

	rows, err := r.db.Query(stmt)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	var users []User
	for rows.Next() {
		var user User

		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.HashedPassword, &user.CreatedAt); err != nil {
			return nil, err
		}

		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return users, nil
}
