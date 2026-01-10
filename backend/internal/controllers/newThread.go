package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

type NewThread struct {
	Topic				string `json:"topic"`
	Description	string `json:"desc"`
}

func HandleNewThread (c *gin.Context) {
	var payload NewThread

	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}
	
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	user, dbErr1 := dataaccess.GetUser(username)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	thread, dbErr2 := dataaccess.CreateThread(payload.Topic, payload.Description, user.ID);
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
	}

	urlCode, urlErr := helpers.EncodeUrl(thread.ID, "thread")
	if urlErr != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
	}

	c.JSON(http.StatusOK, gin.H{
		"urlCode": urlCode,
	})
}
