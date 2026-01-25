// Package models contains the structs that correspond to the rows of the database.
package models

import (
	"time"
)

// Comment is the format of an entry in the "comments" table.
type Comment struct {
	ID        int       `json:"id" gorm:"primary_key"` // Unique identifier for the comment entry.
	Content   string    `json:"content"`               // The content that is written for the comment.
	ThreadID  int       `json:"thread_id"`             // The ID of the thread that the comment is posted in.
	UserID    int       `json:"user_id"`               // The ID of the user who posted the comment.
	CreatedAt time.Time `json:"created_at"`            // The time at which the user posted the comment.
	UpdatedAt time.Time `json:"updated_at"`            // The time at which the comment is last updated.
}
