// Package controllers contains the handler functions for each of the HTTP requests.
package controllers

import (
	"backend/internal/dataaccess"
	"backend/internal/helpers"
	"net/http"

	"github.com/gin-gonic/gin"
)

// HandleViewThread accepts the current gin context of the http request.
// If successful, it updates the context with a HTTP OK response containing the information about the thread
// and a list of comments that have been posted in the thread.
func HandleViewThread(c *gin.Context) {
	// Obtain the urlCode for the thread from the request path.
	urlCode := c.Param("urlCode")

	// Decode the urlCode and verify that it is a valid urlCode for a thread.
	decoded, urlErr := helpers.DecodeUrl(urlCode)
	if urlErr != nil {
		c.JSON(http.StatusNotFound, urlErr.Error())
	}
	if decoded.PageType != "thread" {
		c.JSON(http.StatusNotFound, "Wrong page type for URL code")
	}

	// Obtain information about the thread stored in the threads database
	threadInfo, dbErr3 := dataaccess.GetThreadInfo(decoded.ID)
	if dbErr3 != nil {
		c.JSON(http.StatusInternalServerError, dbErr3.Error())
	}

	// Obtain the username of the host of the thread by id
	host, dbErr4 := dataaccess.GetUserByID(threadInfo.HostID)
	if dbErr4 != nil {
		c.JSON(http.StatusInternalServerError, dbErr4.Error())
	}
	hostname := ""
	if host != nil {
		hostname = host.Username
	}

	// Obtain the ids of the comments posted on this thread by id
	listComments, dbErr5 := dataaccess.GetCommentsFromThread(threadInfo.ID)
	if dbErr5 != nil {
		c.JSON(http.StatusInternalServerError, dbErr5.Error())
	}

	// Generate a comment object for each comment posted on the thread
	threadComments := make([]gin.H, len(listComments))
	for i, cmt := range listComments {
		// Obtain the username of the author of the comment by id
		author, dbErr6 := dataaccess.GetUserByID(cmt.UserID)
		if dbErr6 != nil {
			c.JSON(http.StatusInternalServerError, dbErr6.Error())
		}
		authorname := ""
		if author != nil {
			authorname = author.Username
		}

		// Generate the urlCode for the comment.
		urlCode, urlErr := helpers.EncodeUrl(cmt.ID, "comment")
		if urlErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
		}

		// Obtain the number of likes of the comment.
		numLikes, dbErr7 := dataaccess.CommentLikes(cmt.ID)
		if dbErr7 != nil {
			c.JSON(http.StatusInternalServerError, dbErr7.Error())
		}

		// Add the comment object to the thread comments array.
		threadComments[i] = gin.H{
			"content":   cmt.Content,   // The content of the comment.
			"author":    authorname,    // The username of the author of the comment.
			"urlCode":   urlCode,       // The urlCode of the comment.
			"likes":     numLikes,      // The number of likes received by the comment.
			"createdAt": cmt.CreatedAt, // The time at which the comment is created.
		}
	}

	// Update the context with a HTTP OK response.
	c.JSON(http.StatusOK, gin.H{
		"topic":       threadInfo.Topic,       // The topic title of the thread.
		"description": threadInfo.Description, // The description of the thread.
		"host":        hostname,               // The username of the host of the thread.
		"createdAt":   threadInfo.CreatedAt,   // The time at which the thread is created.
		"comments":    threadComments,         // The list of comments posted in the thread.
	})
}
