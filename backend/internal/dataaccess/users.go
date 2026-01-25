// Package dataaccess contains the functions that creates and retrieves entries from the database.
package dataaccess

import (
	"backend/internal/database"
	"backend/internal/models"
	"errors"

	"gorm.io/gorm"
)

// CreateUser takes in the username and hashed password of the user and creates a new user entry in the database
// It returns the created user object on success, and an error on failure.
func CreateUser(username string, passHash string) (*models.User, error) {
	user := &models.User{
		Username: username,
		Passhash: passHash,
	}
	res := database.DB.Create(user)
	if res.Error != nil {
		return nil, res.Error
	}
	return user, nil
}

// GetUserByName takes in the username of the user and retrieves the relevant user object.
// It returns the retrieved user subject on success, and an error on failure.
func GetUserByName(username string) (*models.User, error) {
	user := &models.User{}
	res := database.DB.Where("username = ?", username).Take(user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return user, nil
}

// GetUserByName takes in the userID of the user and retrieves the relevant user object.
// It returns the retrieved user subject on success, and an error on failure.
func GetUserByID(id int) (*models.User, error) {
	user := &models.User{}
	res := database.DB.Where("id = ?", id).Take(user)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, res.Error
	}
	return user, nil
}
