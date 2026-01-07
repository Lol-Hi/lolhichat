package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
	"os"
	"errors"
	"fmt"
)


type UserClaims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

var key = []byte(os.Getenv("JWT_SECRET_KEY"))

func validateToken[T any, PT interface {*T; jwt.Claims}](tokenStr string) (*T, error) {
	var claims T
	token, err := jwt.ParseWithClaims(tokenStr, PT(&claims), func (token *jwt.Token) (any, error) {
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

func CreateTokens(username string, userID int) (string, string, error) {
	userClaims := UserClaims{
		username,
		jwt.RegisteredClaims{
			Subject: fmt.Sprintf("user-%d", userID),
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer: "lolhichat_admin",
		},
	}

	refreshClaims := jwt.RegisteredClaims{
		Subject: fmt.Sprintf("refresh-%d", userID),
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

	return userTokenStr, refreshTokenStr,  nil
}

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
