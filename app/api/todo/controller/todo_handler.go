package controller

import (
	"net/http"
	"todoBackend/app/api/todo/models"
	"todoBackend/app/api/todo/service"
	"todoBackend/utils/jwts"
	. "todoBackend/utils/responses"

	"github.com/gin-gonic/gin"
)

func AddAllTodo(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "ExtractTokenID failed"))
		return
	}
	var todoList []models.TodoRequest
	if err := c.ShouldBindJSON(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	if err := service.AddAllTodo(todoList, userId); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "service error"))
	}
	c.JSON(http.StatusOK, SuccessResponse(todoList, "Add success!"))

}

// GetAllTodo 获取所有todo
func GetAllTodo(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	todos, err := service.GetAllTodo(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "获取全部todo失败"))
	}
	c.JSON(http.StatusOK, SuccessResponse(todos, "get successfully!"))
}

// 更新现有todo
func UpdateTodoList(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "ExtractTokenID failed"))
		return
	}
	var todoList []models.TodoRequest
	if err := c.ShouldBindJSON(&todoList); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	if err := service.UpdateTodoList(todoList, userId); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "service error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todoList, "Update success!"))
}
