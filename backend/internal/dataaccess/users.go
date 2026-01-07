package dataaccess

import (
	"backend/internal/models"
	"gorm.io/gorm"
	"errors"
)

func CreateUser(username string, passHash string) error {
	user := &models.User{
		Username: username,
		Passhash: passHash,
	}
	res := models.DB.Create(user)
	return res.Error
}

func GetUser(username string) (*models.User, error) {
	user := &models.User{}
	res := models.DB.Where("username = ?", username).Take(user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return user, nil
}

