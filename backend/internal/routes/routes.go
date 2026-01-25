// Package routes contains the functions to link the http routes to the backend handlers.
package routes

import (
	"backend/internal/controllers"
	"backend/internal/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// SetupRouter sets up the router by linking the various middleware and handler functions to the router.
// It returns a gin engine that links the requests from the various http routes to the corresponding handlers.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Origin", "Content-Type", "Accept", "Authorization"}
	router.Use(cors.New(config))

	router.Use(middleware.SetupHeaders())

	router.GET("/viewThread/:urlCode", controllers.HandleViewThread)
	// router.GET("/viewComment/:urlCode", controllers.HandleViewComment)

	router.POST("/search", controllers.HandleSearch)
	router.POST("/login", controllers.HandleLogin)
	router.POST("/signUp", controllers.HandleSignUp)
	router.POST("/renew", controllers.HandleRenewToken)

	protected := router.Group("/api")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/home", controllers.HandleHomePage)
		protected.POST("/newThread", controllers.HandleNewThread)
		protected.POST("/newComment", controllers.HandleNewComment)

		protected.GET("/commentLiked/:urlCode", controllers.HandleCommentLikedStatus)
		protected.POST("/likeComment/:urlCode", controllers.HandleLike)
		protected.POST("/unlikeComment/:urlCode", controllers.HandleUnlike)
	}

	return router
}
