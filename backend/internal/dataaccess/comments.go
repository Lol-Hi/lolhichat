package dataaccess

import (
	"backend/internal/models"
)

func CreateComment(content string, userID int, commentID int) (*models.Comment, error) {
	comment := &models.Comment{
		Content: 	content,
		UserID:		userID,
		ThreadID:	commentID,
	}
	res := models.DB.Create(comment)
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

func GetCommentInfo(commentID int) (*models.Comment, error) {
	comment := &models.Comment{}
	res := models.DB.Where("id = ?", commentID).Take(&comment);
	if res.Error != nil {
		return nil, res.Error
	}
	return comment, nil
}

func GetCommentsFromThread(threadID int) ([]models.Comment, error) {
	var comments []models.Comment
	res := models.DB.Where("thread_id = ?", threadID).Find(&comments);
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}

func SearchComment(query string) ([]models.Comment, error) {
	var comments []models.Comment
	condition := "%" + query + "%"
	res := models.DB.Where("topic LIKE ?", condition).Or("description LIKE ?", condition).Find(&comments)
	if res.Error != nil {
		return nil, res.Error
	}
	return comments, nil
}
