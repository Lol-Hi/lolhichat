// Package dataaccess contains the functions that creates and retrieves entries from the database.
package dataaccess

import (
	"backend/internal/database"
	"backend/internal/models"
	"errors"

	"gorm.io/gorm"
)

// CreateLike takes in the userID and commentID to create a new like relation in the database.
// It returns the created CommentsLike object on success, and an error on failure.
func CreateLike(userID int, commentID int) (*models.CommentsLike, error) {
	like := &models.CommentsLike{
		UserID:    userID,
		CommentID: commentID,
	}
	res := database.DB.Create(like)
	if res.Error != nil {
		return nil, res.Error
	}
	return like, nil
}

// DeleteLike takes in the userID and commentID to search for and delete an existing like relation in the database.
// It returns the deleted CommentsLike object on success, and an error on failure.
func DeleteLike(userID int, commentID int) (*models.CommentsLike, error) {
	like := &models.CommentsLike{
		UserID:    userID,
		CommentID: commentID,
	}
	res := database.DB.Where("comment_id = ?", commentID).Where("user_id = ?", userID).Delete(like)
	if res.Error != nil {
		return nil, res.Error
	}
	return like, nil
}

// CheckUserLike takes in the userID and commentID to check if the likes relation exists in the database.
// It returns true if the like exists, false if the like does not exist, and an error on failure.
func CheckUserLike(userID int, commentID int) (bool, error) {
	like := &models.CommentsLike{}
	res := database.DB.Where("comment_id = ?", commentID).Where("user_id = ?", userID).Take(like)
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

// CommentLikes takes in a commentID and counts the number of likes received by the comment.
// It returns the number of liks associated to the comment on success, and an error on failure.
func CommentLikes(commentID int) (int, error) {
	var numLikes int64
	res := database.DB.Model(&models.CommentsLike{}).Where("comment_id = ?", commentID).Count(&numLikes)
	if res.Error != nil {
		return -1, res.Error
	}
	return int(numLikes), nil
}
