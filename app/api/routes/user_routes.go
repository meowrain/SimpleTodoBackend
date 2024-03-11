package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers"
)

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.POST("/register", handlers.Register)
		userGroup.POST("/login", handlers.Login)
		//userGroup.GET("/:id")
	}
}
