package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/app/models"
	"todoBackend/app/service"
	. "todoBackend/utils"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusOK, ErrorResponse(err, "error"))
		return
	}
	if err := service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todo, "Add success!"))
}

func DeleteTodo(c *gin.Context) {

}
