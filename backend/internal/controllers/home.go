// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleHomePage accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing the username of the user.
func HandleHomePage(c *gin.Context) {
	// Retrieve the username from the context.
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	// Update the context with the HTTP OK response.
	c.JSON(http.StatusOK, gin.H{
		"username": username, // The username of the currently logged-in user.
	})
}
