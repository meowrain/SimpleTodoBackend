package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers"
)

// SetupTodoRoutes 用于配置todo相关的路由
func SetupTodoRoutes(router *gin.Engine) {
	todoGroup := router.Group("/todos")
	{
		// 添加todo
		todoGroup.POST("/add", handlers.CreateTodo)

		//删除todo
		todoGroup.POST("/delete")

		//修改todo
		todoGroup.PUT("/update")

		//获取所有todo
		todoGroup.GET("/all", handlers.GetAllTodo)

		//获取指定id的todo
		todoGroup.GET("/:id")

		//获取todo的条数
		todoGroup.GET("/num")

		//todo图片/语音上传
		todoGroup.POST("/upload")

		//
	}
}
