package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"backend/internal/dataaccess"
	"backend/internal/helpers"
)

type SearchQuery struct {
	Query string `json:"query"`
}

func HandleSearch(c *gin.Context) {
	var payload SearchQuery;

	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	threads, dbErr1 := dataaccess.SearchThread(payload.Query)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}

	results := make([]gin.H, len(threads))
	for i, thd := range threads {
		host, dbErr2 := dataaccess.GetUserByID(thd.HostID)
		if dbErr2 != nil {
			c.JSON(http.StatusInternalServerError, dbErr2.Error())
		}
		hostname := ""
		if host != nil {
			hostname = host.Username
		}

		urlCode, urlErr := helpers.EncodeUrl(thd.ID, "thread")
		if urlErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
		}

		results[i] = gin.H{
			"topic": 				thd.Topic,
			"description": 	thd.Description,
			"host": 				hostname,
			"urlCode":			urlCode,
			"createdAt":		thd.CreatedAt,
		}
	}

	c.JSON(http.StatusOK, results);
	
}
