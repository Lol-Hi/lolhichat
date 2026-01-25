// Package middleware contains the middleware that processes the headers of the HTTP request.
package middleware

import (
	"backend/internal/helpers"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// AuthMiddleware returns a function that takes in the context of the http request and handles the authorization header.
// The function retrieves the JWT user token from the authorization header, and parses it to retrieve the encoded username.
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Missing authorization header"})
			c.Abort()
			return
		}

		headerParts := strings.Split(authHeader, " ")
		if len(headerParts) != 2 || headerParts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Incorrect authorization header format"})
			c.Abort()
			return
		}

		username, jwtErr := helpers.ParseUserToken(headerParts[1])
		if jwtErr != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": jwtErr.Error()})
			c.Abort()
			return
		}

		c.Set(helpers.ContextUsername, username)
		c.Next()
	}
}
