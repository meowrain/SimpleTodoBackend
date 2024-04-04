package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers/feedBackHandler"
)

func SetupFeedBackRoutes(router *gin.Engine) {
	feedBackGroup := router.Group("/feedback")
	{
		feedBackGroup.POST("/add", feedBackHandler.AddFeedback)
		feedBackGroup.GET("/helpful", feedBackHandler.IncrementHelpful)
		feedBackGroup.GET("/helpless", feedBackHandler.IncrementHelpless)
	}
}
