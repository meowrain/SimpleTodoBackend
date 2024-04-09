package feedBackHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/app/api/service/feedBackService"
	"todoBackend/app/models"
	"todoBackend/utils/responses"
)

// AddFeedback 处理添加反馈的请求
func AddFeedback(c *gin.Context) {
	// 绑定 JSON 请求至 comments 变量
	var comments []models.Comment
	if err := c.ShouldBindJSON(&comments); err != nil {
		c.JSON(http.StatusBadRequest, responses.SuccessResponse(comments, "error"))
		return
	}
	// 调用 feedBackService 的 AddComment 方法添加评论
	feedBackService.AddComment(comments)
	c.JSON(http.StatusOK, responses.SuccessResponse(comments, "Feedback and comments added successfully"))
}

// IncrementHelpful 处理增加有帮助反馈计数的请求
func IncrementHelpful(c *gin.Context) {
	// 调用 feedBackService 的 IncrementHelpful 方法增加帮助计数
	err := feedBackService.IncrementHelpful()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse("increment failed", err.Error()))
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("Helpful count incremented", ""))
}

// IncrementHelpless 处理增加无帮助反馈计数的请求
func IncrementHelpless(c *gin.Context) {
	// 调用 feedBackService 的 IncrementHelpLess 方法增加无帮助计数
	err := feedBackService.IncrementHelpLess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse("increment failed", err.Error()))
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("Helpless count incremented", ""))
}
