package dataaccess

import (
	"backend/internal/models"
)

func CreateThread(topic string, description string, hostId int) (*models.Thread, error) {
	thread := &models.Thread{
		Topic:				topic,
		Description:	description,
		HostID:				hostId,
	}
	res := models.DB.Create(thread)
	if res.Error != nil {
		return nil, res.Error
	}
	return thread, nil
}

func GetThreadInfo(threadID int) (*models.Thread, error) {
	thread := &models.Thread{}
	res := models.DB.Where("id = ?", threadID).Take(thread);
	if res.Error != nil {
		return nil, res.Error
	}
	return thread, nil
}

func SearchThread(query string) ([]models.Thread, error) {
	var threads []models.Thread
	condition := "%" + query + "%"
	res := models.DB.Where("topic LIKE ?", condition).Or("description LIKE ?", condition).Find(&threads)
	if res.Error != nil {
		return nil, res.Error
	}
	return threads, nil
}
