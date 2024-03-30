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
		userGroup.Static("/avatars", "./app/static/avatars")
		userGroup.POST("/register", userHandler.Register)
		userGroup.POST("/login", userHandler.Login)
		userGroup.Use(middleware.JwtAuthMiddleware())
		userGroup.GET("/info", userHandler.CurrentUser)
		userGroup.PUT("/update", userHandler.UpdateUser)
		//临时测试
		userGroup.GET("/default_avatar", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"url": "http://127.0.0.1:8090/users/static/user.svg",
			})
		})
		//上传头像
		userGroup.POST("/upload_avatar", userHandler.UploadAvatar)
		userGroup.PUT("/updatebio", userHandler.UpdateBio)
		userGroup.DELETE("/delete", userHandler.DleBio)

	}
}
