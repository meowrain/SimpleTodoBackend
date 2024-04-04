package feedBackHandler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/app/models"
	"todoBackend/app/service/feedBackService"
	"todoBackend/utils/responses"
)

func AddFeedback(c *gin.Context) {
	var comments []models.Comment
	if err := c.ShouldBindJSON(&comments); err != nil {
		c.JSON(http.StatusBadRequest, responses.SuccessResponse(comments, "error"))
		return
	}
	feedBackService.AddComment(comments)
	c.JSON(http.StatusBadRequest, responses.SuccessResponse(comments, "Feedback and comments added successfully"))
}

func IncrementHelpful(c *gin.Context) {
	err := feedBackService.IncrementHelpful()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse("increment failed", err.Error()))
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("Helpful count incremented", ""))
}

func IncrementHelpless(c *gin.Context) {
	err := feedBackService.IncrementHelpLess()
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse("increment failed", err.Error()))
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("Helpless count incremented", ""))
}
