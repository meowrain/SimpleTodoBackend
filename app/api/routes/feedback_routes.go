package routes

import (
	"github.com/gin-gonic/gin"
	"todoBackend/app/api/handlers/feedBackHandler"
)

// SetupFeedBackRoutes 设置反馈路由
func SetupFeedBackRoutes(router *gin.Engine) {
	feedBackGroup := router.Group("/feedback")
	{
		feedBackGroup.POST("/add", feedBackHandler.AddFeedback)           // 添加反馈
		feedBackGroup.GET("/helpful", feedBackHandler.IncrementHelpful)   // 帮助增加
		feedBackGroup.GET("/helpless", feedBackHandler.IncrementHelpless) // 帮助减少
	}
}
