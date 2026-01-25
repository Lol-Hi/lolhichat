// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// SearchQuery is the format of the payload for a search POST request.
type SearchQuery struct {
	Query string `json:"query"` // The query string of the search.
}

// HandleSearch accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing a list of search results.
// Each search result contains some basic thread information.
func HandleSearch(c *gin.Context) {
	// Retrieve the payload from the current context.
	var payload SearchQuery
	reqErr := c.ShouldBindJSON(&payload)
	if reqErr != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": reqErr.Error()})
		return
	}

	// Search for the query in the threads database.
	threads, dbErr1 := dataaccess.SearchThread(payload.Query)
	if dbErr1 != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": dbErr1.Error()})
		return
	}

	// Generate a thread info object for each thread found in the search.
	results := make([]gin.H, len(threads))
	for i, thd := range threads {
		// Get the details of the host of the user.
		host, dbErr2 := dataaccess.GetUserByID(thd.HostID)
		if dbErr2 != nil {
			c.JSON(http.StatusInternalServerError, dbErr2.Error())
		}
		hostname := ""
		if host != nil {
			hostname = host.Username
		}

		// Generate the urlCode for the thread.
		urlCode, urlErr := helpers.EncodeUrl(thd.ID, "thread")
		if urlErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
		}

		// Add the thread information object to the results array.
		results[i] = gin.H{
			"topic":       thd.Topic,       // The topic title of the thread.
			"description": thd.Description, // The description of the thread.
			"host":        hostname,        // The username for the host of the thread.
			"urlCode":     urlCode,         // The urlCode of the thread.
			"createdAt":   thd.CreatedAt,   // The time at which the thread is created.
		}
	}

	// Update the context with a HTTP OK response.
	c.JSON(http.StatusOK, results)

}
