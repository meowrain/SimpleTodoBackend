package main

import (
	"github.com/gin-contrib/cors" // 导入CORS中间件
	"github.com/gin-gonic/gin"    // 导入Gin框架
	"time"
	"todoBackend/app/api/routes" // 导入自定义API路由
	"todoBackend/app/config"     // 导入应用配置
	"todoBackend/utils"          // 导入自定义工具函数
)

func main() {
	utils.ConnectDB()       // 连接数据库
	utils.CreateTable()     // 创建数据库表
	router := gin.Default() // 创建Gin框架实例

	router.Use(cors.New(cors.Config{ // 使用CORS中间件
		AllowAllOrigins:  true,                                                                  // 允许所有来源
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},          // 允许的HTTP方法
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                                            // 公开的响应头
		AllowCredentials: true,                                                                  // 允许发送凭据
		MaxAge:           12 * time.Hour,                                                        // 预检请求的有效期
	}))

	// 设置路由
	routes.SetupUserRoutes(router)     // 设置用户相关路由
	routes.SetupTodoRoutes(router)     // 设置待办事项相关路由
	routes.SetupFeedBackRoutes(router) // 设置反馈相关路由

	router.Run(":" + config.Cfg.Server.AppPort) // 运行服务并指定端口
}
