package routes

import "github.com/gin-gonic/gin"

func SetupUserRoutes(router *gin.Engine) {
	userGroup := router.Group("/users")
	{
		userGroup.GET("/", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"msg": "Test",
			})
		})

		//userGroup.GET("/:id")
	}
}
