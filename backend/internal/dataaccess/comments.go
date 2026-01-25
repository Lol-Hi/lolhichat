// Package dataaccess contains the functions that creates and retrieves entries from the database.
package dataaccess

import (
	"backend/internal/database"
	"backend/internal/models"
)

// CreateComment takes in the content, userID and threadsID to create a new comment entry in the database.
// It returns the created comment object on success, and an error on failure.
func CreateComment(content string, userID int, threadID int) (*models.Comment, error) {
	comment := &models.Comment{
		Content:  content,
		UserID:   userID,
		ThreadID: threadID,
	}
	res := database.DB.Create(comment)
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

// GetCommentInfo takes in the commentID to retrieve the relevant comment entry from the database.
// It returns the retrieved comment object on success, and an error on failure.
func GetCommentInfo(commentID int) (*models.Comment, error) {
	comment := &models.Comment{}
	res := database.DB.Where("id = ?", commentID).Take(&comment)
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

// GetCommentsFromThread takes in the threadID and searches for all comments in that thread.
// It returns a slice of comments on success, and an error on failure.
func GetCommentsFromThread(threadID int) ([]models.Comment, error) {
	var comments []models.Comment
	res := database.DB.Where("thread_id = ?", threadID).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
