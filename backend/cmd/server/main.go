package main

import (
	"backend/internal/database"
	"backend/internal/helpers"
	"backend/internal/routes"
	"log"

	"github.com/joho/godotenv"
)

func main() {
	envErr := godotenv.Load()
	if envErr != nil {
		log.Fatal("Error loading .env file")
	}

	dbErr := database.ConnectDatabase()
	if dbErr != nil {
		log.Fatal("Error connecting to database")
	}

	squidErr := helpers.InitSquid()
	if squidErr != nil {
		log.Fatal("Error initialising sqids")
	}

	router := routes.SetupRouter()
	router.Run("localhost:8080")
}
