package helpers

import (
	"golang.org/x/crypto/bcrypt"
	"errors"
)

func HashPassword(password string, cost int) (string, error) {
  bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	return string(bytes), err
}
 
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
