// Package middleware contains the middleware that processes the headers of the HTTP request.
package middleware

import (
	"github.com/gin-gonic/gin"
)

func SetupHeaders() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
	}
}
