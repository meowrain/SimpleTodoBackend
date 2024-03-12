package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers"
	"todoBackend/app/middleware"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", handlers.Register)
		userGroup.POST("/login", handlers.Login)
		userGroup.Use(middleware.JwtAuthMiddleware())
		userGroup.GET("/", handlers.CurrentUser)
		userGroup.PUT("/update", handlers.UpdateUser)
	}
}
