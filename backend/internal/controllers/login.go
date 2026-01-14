package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func HandleLogin (c *gin.Context) {
	var payload Login

	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	user, dbErr := dataaccess.GetUserByName(payload.Username)
	if dbErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	isMatch, hashErr := helpers.VerifyPassword(payload.Password, user.Passhash)
	if hashErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": hashErr.Error()})
		return
	}
	if !isMatch {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Password"})
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
