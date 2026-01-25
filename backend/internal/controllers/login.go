// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Login is the format of the payload from a login POST request.
type Login struct {
	Username string `json:"username"` // The username entered for the login request.
	Password string `json:"password"` // The password entered for the login request.
}

// HandleLogin accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing the current user and refresh tokens
func HandleLogin(c *gin.Context) {
	// Retrieve the payload form the current context
	var payload Login
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Obtain user data based on the username
	user, dbErr := dataaccess.GetUserByName(payload.Username)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	// Verify that the password given in the payload corresponds with the hashed password in the database
	isMatch, hashErr := helpers.VerifyPassword(payload.Password, user.Passhash)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
		return
	}
	if !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
		return
	}

	// Generate user and refresh tokens
	userToken, refreshToken, tokenErr := helpers.CreateTokens(user.Username, user.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	// Update the context with the HTTP OK response
	c.JSON(http.StatusOK, gin.H{
		"userToken":    userToken,    // The JWT user token for the user session.
		"refreshToken": refreshToken, // The JWT refresh token to renew the user token after it expires.
	})
}
