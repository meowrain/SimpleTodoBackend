package routes

import "github.com/gin-gonic/gin"

func SetupTodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.GET("/")
		todoGroup.POST("/")
	}
}
