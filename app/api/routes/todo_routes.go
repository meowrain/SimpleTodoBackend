package routes

import "github.com/gin-gonic/gin"

// SetupTodoRoutes 用于配置todo相关的路由
func SetupTodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/")
		todoGroup.POST("/")
	}
}
