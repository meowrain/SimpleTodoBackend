package main

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/routes"
	"todoBackend/app/config"
	"todoBackend/app/middleware"
	"todoBackend/utils"
)

func main() {

	utils.CreateTable()
	router := gin.Default()
	router.Use(middleware.ConnectDBMiddleWare())
	routes.SetupUserRoutes(router)
	routes.SetupTodoRoutes(router)
	router.Run(":" + config.Cfg.Server.AppPort)

}
