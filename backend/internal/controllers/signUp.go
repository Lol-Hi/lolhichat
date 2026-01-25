// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// SignUp is the format of the payload for a sign-up POST request.
type SignUp struct {
	Username string `json:"username"` // Username of the newly registered user.
	Password string `json:"password"` // Password of the newly registered user.
}

// A constant for the hash function for the password.
const HASH_COST = 14

// HandleSearch accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing a list of search results.
func HandleSignUp(c *gin.Context) {
	// Retrieve the payload from the currrent context
	var payload SignUp
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Check if there is already an existing user with the same username
	existingUser, dbErr1 := dataaccess.GetUserByName(payload.Username)
	if existingUser != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "Username already exists"})
		return
	}
	if dbErr1 != nil && !errors.Is(dbErr1, gorm.ErrRecordNotFound) {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
	}

	// Hash the password
	passHash, hashErr := helpers.HashPassword(payload.Password, HASH_COST)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
		return
	}

	// Create a new user row in the users database
	_, dbErr2 := dataaccess.CreateUser(payload.Username, passHash)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}

	// Update the context with a HTTP OK response
	c.JSON(http.StatusOK, nil)
}
