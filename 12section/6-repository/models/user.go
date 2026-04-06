package models

import "time"

// User represents a customer in the database with a backing table of users
type User struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Email          string    `json:"email"`
	HashedPassword string    `json:"-"`
	CreatedAt      time.Time `json:"created_at"`
}

// Profile belongs to a user
type Profile struct {
	UserID    int       `json:"id"`
	Avatar    string    `json:"avatar"`
	CreatedAt time.Time `json:"created_at"`
}
