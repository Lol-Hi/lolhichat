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

	threadInfo, dbErr3 := dataaccess.GetThreadInfo(decoded.ID)
	if dbErr3 != nil {
		c.JSON(http.StatusInternalServerError, dbErr3.Error())
	}

	host, dbErr4 := dataaccess.GetUserByID(threadInfo.HostID)
	if dbErr4 != nil {
		c.JSON(http.StatusInternalServerError, dbErr4.Error())
	}
	hostname := ""
	if host != nil {
		hostname = host.Username
	}

	listComments, dbErr5 := dataaccess.GetCommentsFromThread(threadInfo.ID)
	if dbErr5 != nil {
		c.JSON(http.StatusInternalServerError, dbErr5.Error())
	}

	threadComments := make([]gin.H, len(listComments))
	for i, cmt := range listComments {
		author, dbErr6 := dataaccess.GetUserByID(cmt.UserID)
		if dbErr6 != nil {
			c.JSON(http.StatusInternalServerError, dbErr6.Error())
		}

		authorname := "";
		if author != nil {
			authorname = author.Username
		}

		urlCode, urlErr := helpers.EncodeUrl(cmt.ID, "comment")
		if urlErr != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": urlErr.Error()})
		}
		
		numLikes, dbErr7 := dataaccess.CommentLikes(cmt.ID);
		if dbErr7 != nil {
			c.JSON(http.StatusInternalServerError, dbErr7.Error())
		}


		threadComments[i] = gin.H{
			"content": 		cmt.Content,
			"author":			authorname,
			"urlCode":		urlCode,
			"likes":			numLikes,
			"createdAt":	cmt.CreatedAt,
		}
	}

 	c.JSON(http.StatusOK, gin.H{
		"topic": 				threadInfo.Topic,
		"description":	threadInfo.Description,
		"host":					hostname,									
		"createdAt":		threadInfo.CreatedAt,
		"comments":			threadComments,
	})
}
