package main

import (
	"github.com/joho/godotenv"
	"log"
	"backend/internal/models"
	"backend/internal/routes"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbErr := models.ConnectDatabase()
	if dbErr != nil {
		log.Fatal("Error connecting to database")
	}

	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
