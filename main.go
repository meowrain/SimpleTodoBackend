package main

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/routes"
	"todoBackend/app/config"
	"todoBackend/utils"
)

func main() {
	utils.ConnectDB()
	utils.CreateTable()
	router := gin.Default()

	// 设置路由
	routes.SetupUserRoutes(router)
	routes.SetupTodoRoutes(router)

	router.Run(":" + config.Cfg.Server.AppPort)
}
