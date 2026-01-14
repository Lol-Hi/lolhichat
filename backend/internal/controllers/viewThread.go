package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/helpers"
	"backend/internal/dataaccess"
)

func HandleViewThread(c *gin.Context) {
	urlCode := c.Param("urlCode")
	
	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "thread" {
		c.JSON(http.StatusNotFound, "Wrong page type for URL code")
	}

	threadInfo, dbErr1 := dataaccess.GetThreadInfo(decoded.ID)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, dbErr1.Error())
	}

	host, dbErr2 := dataaccess.GetUserByID(threadInfo.HostID)
	if dbErr2 != nil {
		c.JSON(http.StatusInternalServerError, dbErr2.Error())
	}
	hostname := ""
	if host != nil {
		hostname = host.Username
	}

 	c.JSON(http.StatusOK, gin.H{
		"topic": 				threadInfo.Topic,
		"description":	threadInfo.Description,
		"host":					hostname,									
		"createdAt":		threadInfo.CreatedAt,
	})
}
