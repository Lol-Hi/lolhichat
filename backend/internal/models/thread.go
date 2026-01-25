// Package models contains the structs that correspond to the rows of the database.
package models

import (
	"time"
)

// Thread is the format of an entry in the "threads" table.
type Thread struct {
	ID          int       `json:"id" gorm:"primary_key"` // Unique identifier for the thread entry.
	Topic       string    `json:"topic"`                 // The topic title of the thread.
	Description string    `json:"description"`           // The description of the thread.
	HostID      int       `json:"host_id"`               // The ID of the host who created the thread.
	CreatedAt   time.Time `json:"created_at"`            // The time at which the thread is created.
	UpdatedAt   time.Time `json:"updated_at"`            // The time at which the thread is last updated.
}
