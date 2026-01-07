package middleware

import (
	"net/http"
	"strings"
	"github.com/gin-gonic/gin"
	"backend/internal/helpers"
)

func AuthMiddleware() gin.HandlerFunc {
	return func (c *gin.Context) {
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

