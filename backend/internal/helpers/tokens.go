package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

var key = []byte("sdvaom")

func CreateToken(username string) (string, error) {
	token := jwt.NewWithClaims(
		jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp": time.Now().Add(time.Hour * 24).Unix(),
		})

	tokenStr, err := token.SignedString(key)
	if err != nil {
		return "", err
	}

	return tokenStr, err
}


