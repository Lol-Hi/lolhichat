// Package helpers contains the helper functions for other services used by the api.
// This file contains the functions related to the creation and processing of JWT tokens.
package helpers

import (
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

// UserClaims is the format of the claims object that will be encoded by the JWT string
type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// Obtain the secret key from the .env file
var key = []byte(os.Getenv("JWT_SECRET_KEY"))

// validateToken takes in a JWT token as a string, along with its type and pointer type.
// It checks that the JWT token is valid and decodes the JWT string to retrieve the encoded claims.
// It returns the claims encoded by the token on success, and an error on failure.
// If the token has expired, the claims will still be returned in order to retrieve the user information for token renewal.
func validateToken[T any, PT interface {
	*T
	jwt.Claims
}](tokenStr string) (*T, error) {
	var claims T
	token, err := jwt.ParseWithClaims(tokenStr, PT(&claims), func(token *jwt.Token) (any, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New(fmt.Sprintf("Unexpected signing method: %s", token.Header["alg"]))
		}
		return key, nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return &claims, err
		}
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("Invalid token")
	}

	return &claims, nil
}

// CreateTokens takes in the username and userID of the user to create the JWT user and refresh tokens.
// It returns the created user and refresh tokens on success, and an error on failure.
func CreateTokens(username string, userID int) (string, string, error) {
	userClaims := UserClaims{
		username,
		jwt.RegisteredClaims{
			Subject:   fmt.Sprintf("user-%d", userID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "lolhichat_admin",
		},
	}

	refreshClaims := jwt.RegisteredClaims{
		Subject:   fmt.Sprintf("refresh-%d", userID),
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
	}

	sign := jwt.SigningMethodHS256
	userToken := jwt.NewWithClaims(sign, userClaims)
	refreshToken := jwt.NewWithClaims(sign, refreshClaims)

	userTokenStr, utErr := userToken.SignedString(key)
	if utErr != nil {
		return "", "", utErr
	}

	refreshTokenStr, rtErr := refreshToken.SignedString(key)
	if rtErr != nil {
		return "", "", rtErr
	}

	return userTokenStr, refreshTokenStr, nil
}

// ParseUserToken takes in a JWT user token and decodes the JWT token to obtain the username.
// It returns the encoded username on success, and an error on failure.
func ParseUserToken(userTokenStr string) (string, error) {
	userClaims, err := validateToken[UserClaims, *UserClaims](userTokenStr)
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", errors.New("Expired User Token")
		}
		return "", err
	}

	return userClaims.Username, nil
}

// ValidateRenew takes in the JWT user and refresh tokens, and checks that the refresh token is valid for the renewal.
// It returns the encoded username on success, and an error on failure.
func ValidateRenew(userTokenStr string, refreshTokenStr string) (string, error) {
	_, refErr := validateToken[jwt.RegisteredClaims, *jwt.RegisteredClaims](refreshTokenStr)
	if refErr != nil {
		return "", refErr
	}

	userClaims, userErr := validateToken[UserClaims, *UserClaims](userTokenStr)
	if userErr != nil && !errors.Is(userErr, jwt.ErrTokenExpired) {
		return "", userErr
	}

	return userClaims.Username, nil
}
