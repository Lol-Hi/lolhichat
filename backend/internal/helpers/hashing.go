// Package helpers contains the helper functions for other services used by the api.
// This file contains the functions related to the hashing of passwords
package helpers

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword takes in a password string and an integer cost.
// It returns the hashed password on success and an error on failure.
func HashPassword(password string, cost int) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}

// VerifyPassword takes in a password attempt and the hashed password from the database.
// It returns the true if the password is correct, false if the password is wrong, and an error on failure during the check.
func VerifyPassword(password string, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		if errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
			return false, nil
		} else {
			return false, err
		}
	}
	return true, nil
}
