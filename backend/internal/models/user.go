package models

type User struct {
	ID				int			`json:"id" gorm:"primary_key"`
	Username	string	`json:"username"`
	Passhash	string	`json:"passHash"`
}
