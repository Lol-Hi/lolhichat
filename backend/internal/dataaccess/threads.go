// Package dataaccess contains the functions that creates and retrieves entries from the database.
package dataaccess

import (
	"backend/internal/database"
	"backend/internal/models"
)

// CreateThread takes in the topic, description and hostID of the thread and creates a new thread entry in the database.
// It returns the created thread object on success, and an error on failure.
func CreateThread(topic string, description string, hostId int) (*models.Thread, error) {
	thread := &models.Thread{
		Topic:       topic,
		Description: description,
		HostID:      hostId,
	}
	res := database.DB.Create(thread)
	if res.Error != nil {
		return nil, res.Error
	}
	return thread, nil
}

// GetThreadInfo takes in a threadID and retrieves the relevant thread entry in the database.
// It returns the retrieved thread object on success, and an error on failure.
func GetThreadInfo(threadID int) (*models.Thread, error) {
	thread := &models.Thread{}
	res := database.DB.Where("id = ?", threadID).Take(thread)
	if res.Error != nil {
		return nil, res.Error
	}
	return thread, nil
}

// SearchThread takes in a search query and searches for all the threads whose topic or description matches the query.
// It returns a slice of all the threads that matches the search query on success, and an error on failure.
func SearchThread(query string) ([]models.Thread, error) {
	var threads []models.Thread
	condition := "%" + query + "%"
	res := database.DB.Where("topic LIKE ?", condition).Or("description LIKE ?", condition).Find(&threads)
	if res.Error != nil {
		return nil, res.Error
	}
	return threads, nil
}
