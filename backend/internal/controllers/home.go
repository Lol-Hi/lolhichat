package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"backend/internal/helpers"
)

func HandleHomePage(c *gin.Context) {
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	c.JSON(http.StatusOK, gin.H{"username": username})
}
