// Package models contains the structs that correspond to the rows of the database.
package models

// User is the format of an entry in the "users" table.
type User struct {
	ID       int    `json:"id" gorm:"primary_key"` // Unique identifier for the user entry.
	Username string `json:"username"`              // The username of the user.
	Passhash string `json:"passHash"`              // The hashed password of the user.
}
