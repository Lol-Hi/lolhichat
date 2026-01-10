package models

import (
	"time"
)

type Comment struct {
	ID				int				`json:"id" gorm:"primary_key"`
	Content 	string		`json:"content"`
	ThreadID	int				`json:"thread_id"`
	UserID		int 			`json:"user_id"`
	ReplyToID int				`json:"reply_to_id"`
	CreatedAt	time.Time	`json:"created_at"`
	UpdatedAt time.Time	`json:"updated_at"`
}
