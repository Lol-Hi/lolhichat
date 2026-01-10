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
