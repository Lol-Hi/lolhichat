package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"backend/internal/handlers/controllers"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	router.Use(cors.New(config))

	router.POST("/login", controllers.HandleLogin)  
	router.POST("/signUp", controllers.HandleSignUp)

	return router
}
