package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers"
	"todoBackend/app/middleware"
)

// SetupUserRoutes 函数负责配置与用户相关的路由。
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
