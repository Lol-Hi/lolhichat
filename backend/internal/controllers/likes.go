// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleCommentLikedStatus accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing a boolean value of the liked status of the post.
func HandleCommentLikedStatus(c *gin.Context) {
	// Obtain the urlCode of the comment from the request path
	urlCode := c.Param("urlCode")

	// Decode the urlCode and verify that it is a valid urlCode for a comment.
	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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

	// Verify if the user has liked the post based on user ID and post ID.
	isLiked, dbErr2 := dataaccess.CheckUserLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}

	// Update the context with the HTTP OK response
	c.JSON(http.StatusOK, gin.H{
		"isLiked": isLiked, // Boolean response to indicate if the user has liked the comment.
	})
}

// HandleLike accepts the current gin context of the http request.
// It updates the likes database with a new like by the specified user.
// If successful, it returns an HTTP OK response.
func HandleLike(c *gin.Context) {
	// Obtain the urlCode of the comment from the request path.
	urlCode := c.Param("urlCode")

	// Decode the urlCode and verify that it is a valid urlCode for a comment.
	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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

	// Create a new like relation in the likes database.
	_, dbErr2 := dataaccess.CreateLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}

	// Update the context witha a HTTP OK response.
	c.JSON(http.StatusOK, gin.H{})
}

// HandleUnLike accepts the current gin context of the http request.
// It deletes the like by the user from the likes database.
// If successful, it returns an HTTP OK response.
func HandleUnlike(c *gin.Context) {
	// Obtain the urlCode of the comment from the request path.
	urlCode := c.Param("urlCode")

	// Decode the urlCode and verify that it is a valid urlCode for a comment.
	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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

	// Remove the relevant like relation in the likes database.
	_, dbErr2 := dataaccess.DeleteLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}

	// Update the context with a HTTP OK response.
	c.JSON(http.StatusOK, gin.H{})
}
