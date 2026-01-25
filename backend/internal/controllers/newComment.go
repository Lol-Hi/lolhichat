// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// NewComment is the format of the payload from a new comment POST request.
type NewComment struct {
	Comment string `json:"comment"` // The content of the comment.
	UrlCode string `json:"urlCode"` // The url code of the thread that the comment is to be posted in.
}

// HandleNewComment accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response.
func HandleNewComment(c *gin.Context) {
	// Retrieve the payload from the current context.
	var payload NewComment
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Decode the urlCode and verify that it is a valid urlCode for a thread.
	decoded, urlErr := helpers.DecodeUrl(payload.UrlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "thread" {
		c.JSON(http.StatusNotFound, "Wrong page type for URL code")
	}

	// Retrieve the username from the context.
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	// Obtain user data based on the username.
	user, dbErr1 := dataaccess.GetUserByName(username)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	// Create a new comment entry in the comments database.
	_, dbErr2 := dataaccess.CreateComment(payload.Comment, user.ID, decoded.ID)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
	}

	// Update the context with the HTTP OK response.
	c.JSON(http.StatusOK, gin.H{})
}
