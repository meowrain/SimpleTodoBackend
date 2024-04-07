package routes

import (
	"todoBackend/app/api/handlers/userHandler"
	"todoBackend/app/middleware"

	"github.com/gin-gonic/gin"
)

// SetupUserRoutes 函数负责配置与用户相关的路由。
func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		//为user设置静态目录，存放用户头像信息
		userGroup.Static("/avatars", "./app/static/avatars") // 设置静态目录用于存放用户头像信息
		userGroup.POST("/register", userHandler.Register)    // 处理用户注册请求
		userGroup.POST("/login", userHandler.Login)          // 处理用户登录请求
		userGroup.Use(middleware.JwtAuthMiddleware())        // 使用JWT认证中间件
		userGroup.GET("/info", userHandler.CurrentUser)      // 获取当前用户信息
		userGroup.PUT("/update", userHandler.UpdateUser)     // 更新用户信息
		//临时测试
		userGroup.GET("/default_avatar", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"url": "http://127.0.0.1:8090/users/static/user.svg",
			})
		})
		//上传头像
		userGroup.POST("/upload_avatar", userHandler.UploadAvatar) // 处理上传用户头像请求

	}
}
