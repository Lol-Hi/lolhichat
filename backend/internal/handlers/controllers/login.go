package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"errors"
	"gorm.io/gorm"
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

	user, dbErr := dataaccess.GetUser(payload.Username)
	if dbErr != nil {
		if errors.Is(dbErr, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Username not found"})
			return
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr.Error()})
			return
		}
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

	token, tokenErr := helpers.CreateToken(payload.Username)
	if tokenErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": tokenErr.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
