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
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "提取用户ID失败"))
		return
	}
	todos, err := service.GetAllTodo(userId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse(err, "获取todo列表失败"))
		return
	}
	count := service.Num_TodoList(todos)
	responseData := map[string]interface{}{
		"todos":   todos,
		"count":   count,
		"message": "获取成功",
	}
	c.JSON(http.StatusOK, SuccessResponse(responseData, "获取成功"))
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
