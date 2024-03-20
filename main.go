package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
	"todoBackend/app/api/routes"
	"todoBackend/app/config"
	"todoBackend/utils"
)

func main() {
	utils.ConnectDB()
	utils.CreateTable()
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	// 设置路由
	routes.SetupUserRoutes(router)
	routes.SetupTodoRoutes(router)

	router.Run(":" + config.Cfg.Server.AppPort)
}
