package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

type NewComment struct {
	Comment	string `json:"comment"`
	UrlCode	string `json:"urlCode"`
}

func HandleNewComment (c *gin.Context) {
	var payload NewComment

	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}
	
	decoded, urlErr := helpers.DecodeUrl(payload.UrlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "thread" {
		c.JSON(http.StatusNotFound, "Wrong page type for URL code")
	}
	
	username, userErr := helpers.GetContextUsername(c)
	if userErr != nil {
		c.JSON(http.StatusInternalServerError, userErr.Error())
	}

	user, dbErr1 := dataaccess.GetUserByName(username)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}
	if user == nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid Username"})
		return
	}

	_, dbErr2 := dataaccess.CreateComment(payload.Comment, user.ID, decoded.ID);
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
	}


	c.JSON(http.StatusOK, gin.H{})
}

