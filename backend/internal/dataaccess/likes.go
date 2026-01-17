package dataaccess

import (
	"gorm.io/gorm"
	"errors"
	"backend/internal/models"
)

func CreateLike(userID int, commentID int) (*models.CommentsLike, error) {
	like := &models.CommentsLike{
		UserID:			userID,
		CommentID:	commentID,
	}
	res := models.DB.Create(like)
	if res.Error != nil {
		return nil, res.Error
	}
	return like, nil
}

func DeleteLike(userID int, commentID int) (*models.CommentsLike, error) {
	like := &models.CommentsLike{
		UserID:			userID,
		CommentID:	commentID,
	}
	res := models.DB.Where("comment_id = ?", commentID).Where("user_id = ?", userID).Delete(like)
	if res.Error != nil {
		return nil, res.Error
	}
	return like, nil
}

func CheckUserLike(userID int, commentID int) (bool, error) {
	like := &models.CommentsLike{}
	res := models.DB.Where("comment_id = ?", commentID).Where("user_id = ?", userID).Take(like);
	if res.Error != nil {
		if errors.Is(res.Error, gorm.ErrRecordNotFound) {
			return false, nil
		}
		return false, res.Error
	}
	return true, nil
}

func CommentLikes(commentID int) (int, error) {
	var numLikes int64
	res := models.DB.Model(&models.CommentsLike{}).Where("comment_id = ?", commentID).Count(&numLikes)
	if res.Error != nil {
		return -1, res.Error
	}
	return int(numLikes), nil;
}
