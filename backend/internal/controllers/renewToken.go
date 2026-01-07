package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

type Renew struct {
	UserToken string `json:"userToken"` 
	RefreshToken string `json:"refreshToken"`
}

func HandleRenewToken (c *gin.Context) {
	var payload Renew

	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}
	
	username, authErr := helpers.ValidateRenew(payload.UserToken, payload.RefreshToken)
	if authErr != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": authErr.Error()})
		return
	}

	user, dbErr := dataaccess.GetUser(username)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Token does not correspond with a valid username"})
		return
	}

	userToken, refreshToken, tokenErr := helpers.CreateTokens(user.Username, user.ID)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"userToken": userToken,
		"refreshToken": refreshToken,
	})
}
