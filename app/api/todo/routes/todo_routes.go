package routes

import (
	"todoBackend/app/api/todo/controller"
	"todoBackend/app/middleware"

	"github.com/gin-gonic/gin"
)

// SetupTodoRoutes 用于配置todo相关的路由
func SetupTodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.Use(middleware.JwtAuthMiddleware())
		// 添加todo
		todoGroup.POST("/all", controller.AddAllTodo)

		// 获取所有todo
		todoGroup.GET("/all", controller.GetAllTodo)
		//更新Todo
		todoGroup.PUT("/all", controller.UpdateTodoList)

	}
}
