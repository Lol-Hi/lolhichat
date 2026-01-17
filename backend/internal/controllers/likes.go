package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

func HandleCommentLikedStatus(c *gin.Context) {
	urlCode := c.Param("urlCode")

	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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
	
	isLiked, dbErr2 := dataaccess.CheckUserLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{"isLiked": isLiked})
}

func HandleLike(c *gin.Context) {
	urlCode := c.Param("urlCode")

	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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
	
	_, dbErr2 := dataaccess.CreateLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{})
}

func HandleUnlike(c *gin.Context) {
	urlCode := c.Param("urlCode")

	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "comment" {
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
	
	_, dbErr2 := dataaccess.DeleteLike(user.ID, decoded.ID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr2.Error()})
		return
	}
	
	c.JSON(http.StatusOK, gin.H{})
}
