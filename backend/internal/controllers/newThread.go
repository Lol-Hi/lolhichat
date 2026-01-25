// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewThread is the format of the payload from a new thread POST request.
type NewThread struct {
	Topic       string `json:"topic"` // The topic title of the new thread.
	Description string `json:"desc"`  // The description of the new thread.
}

// HandleNewThread accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing the urlCode for the thread.
func HandleNewThread(c *gin.Context) {
	// Retrieve the payload from the current context.
	var payload NewThread
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Retrieve the username from the context.
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	// Obtain user data based on username.
	user, dbErr1 := dataaccess.GetUserByName(username)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	// Create a new thread entry in the thread database.
	thread, dbErr2 := dataaccess.CreateThread(payload.Topic, payload.Description, user.ID)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
	}

	// Create a urlCode based on the id of thread entry.
	urlCode, urlErr := helpers.EncodeUrl(thread.ID, "thread")
	if urlErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
	}

	// Update the context with a HTTP OK respose.
	c.JSON(http.StatusOK, gin.H{
		"urlCode": urlCode, // The urlCode of the newly created thread.
	})
}
