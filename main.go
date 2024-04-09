package main

import (
	"time"
	"todoBackend/app/api/routes" // 导入自定义API路由
	"todoBackend/app/config"     // 导入应用配置
	"todoBackend/utils"          // 导入自定义工具函数

	_ "todoBackend/docs" // 导入Swagger文档

	"github.com/gin-contrib/cors" // 导入CORS中间件
	"github.com/gin-gonic/gin"    // 导入Gin框架
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title           SimpleTodo后端接口文档
// @version         1.0
// @description     SimpleTodo后端接口文档
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8090
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/
func main() {
	utils.ConnectDB()   // 连接数据库
	utils.CreateTable() // 创建数据库表
	gin.SetMode(gin.ReleaseMode)
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
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":" + config.Cfg.Server.AppPort) // 运行服务并指定端口
}
