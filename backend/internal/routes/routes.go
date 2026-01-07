package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"backend/internal/controllers"
	"backend/internal/middleware"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))
	
	router.Use(middleware.SetupHeaders())
	
	router.POST("/login", controllers.HandleLogin)  
	router.POST("/signUp", controllers.HandleSignUp)
	router.POST("/renew", controllers.HandleRenewToken)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/home", controllers.HandleHomePage)
	}

	return router
}
