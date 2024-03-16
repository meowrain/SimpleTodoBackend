package main

import (
	"todoBackend/app/api/routes"
	"todoBackend/app/config"
	"todoBackend/app/middleware"
	"todoBackend/utils"

	"github.com/gin-gonic/gin"
)

func main() {
	utils.CreateTable()
	router := gin.Default()
	router.Use(middleware.ConnectDBMiddleWare())

	// 设置路由
	routes.SetupUserRoutes(router)
	routes.SetupTodoRoutes(router)

	router.Run(":" + config.Cfg.Server.AppPort)
}
