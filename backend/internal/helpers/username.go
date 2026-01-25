// Package helpers contains the helper functions for other services used by the api.
// This file contains the functions related to obtaining the username from the protected context.
package helpers

import (
	"errors"

	"github.com/gin-gonic/gin"
)

// A constant to store the username of the currently logged-in user.
const ContextUsername = ""

// GetContextUsername takes in the current gin context and retrieves the username in the protected context.
// It returns the username on success, and an error on failure.
func GetContextUsername(c *gin.Context) (string, error) {
	user, exists := c.Get(ContextUsername)
	if !exists {
		return "", errors.New("No username in context")
	}
	username, ok := user.(string)
	if !ok {
		return "", errors.New("Context username of invalid type")
	}

	return username, nil
}
