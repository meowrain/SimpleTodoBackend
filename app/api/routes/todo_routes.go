package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/controller/todoHandler"
	"todoBackend/app/middleware"
)

// SetupTodoRoutes 用于配置todo相关的路由
func SetupTodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		todoGroup.Use(middleware.JwtAuthMiddleware())
		// 添加todo
		todoGroup.POST("/add", todoHandler.CreateTodo)

		// 删除todo
		todoGroup.DELETE("/delete/:id", todoHandler.DeleteTodo)

		// 修改todo
		todoGroup.PUT("/update/:id", todoHandler.UpdateTodo)

		// 获取所有todo
		todoGroup.GET("/all", todoHandler.GetAllTodo)

		// 获取指定id的todo
		todoGroup.GET("/:id", todoHandler.GetTodo)

		// 获取todo的条数
		todoGroup.GET("/num", todoHandler.GetNumofTodo)

		// todo图片/语音上传
		todoGroup.POST("/upload/:id", todoHandler.UploadTodoPhoto)

		//
	}
}
