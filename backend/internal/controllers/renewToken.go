// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Renew is the format of the payload from a token renewal POST request.
type Renew struct {
	UserToken    string `json:"userToken"`    // The existing JWT user token.
	RefreshToken string `json:"refreshToken"` // The existing JWT refresh token.
}

// HandleRenewToken accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing the new user and refresh tokens.
func HandleRenewToken(c *gin.Context) {
	// Retrieve the payload from the current context.
	var payload Renew
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Validate the current user and refresh tokens and obtain the username
	username, authErr := helpers.ValidateRenew(payload.UserToken, payload.RefreshToken)
	if authErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": authErr.Error()})
		return
	}

	// Obtain user data based on username
	user, dbErr := dataaccess.GetUserByName(username)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token does not correspond with a valid username"})
		return
	}

	// Generate a new pair of user and refresh tokens from the username and user id
	userToken, refreshToken, tokenErr := helpers.CreateTokens(user.Username, user.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	// Update the context with a HTTP OK response
	c.JSON(http.StatusOK, gin.H{
		"userToken":    userToken,    // The renewed JWT user token.
		"refreshToken": refreshToken, // The renewed JWT refresh token.
	})
}
