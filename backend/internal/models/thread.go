package models

import (
	"time"
)

type Thread struct {
	ID					int				`json:"id" gorm:"primary_key"`
	Topic				string		`json:"topic"`
	Description	string		`json:"description"`
	HostID			int				`json:"host_id"`
	CreatedAt		time.Time	`json:"created_at"`
	UpdatedAt		time.Time `json:"updated_at"`
}
