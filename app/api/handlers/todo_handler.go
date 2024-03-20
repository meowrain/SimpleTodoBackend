package handlers

import (
	"net/http"
	"strconv"
	"todoBackend/app/models"
	"todoBackend/app/service"
	. "todoBackend/utils"

	"github.com/gin-gonic/gin"
)

func CreateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	if err := service.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todo, "Add success!"))
}
func GetAllTodo(c *gin.Context) {
	todos, err := service.GetAllTodo()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "获取全部todo失败"))
	}
	c.JSON(http.StatusOK, SuccessResponse(todos, "get successfully!"))
}

func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.DeleteTodo(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "删除todo失败"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(id, "todo 删除成功"))
}

func UpdateTodo(c *gin.Context) {
	var todo models.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := service.UpdateTodo(id, &todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todo, "Update success!"))
}
func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := service.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(err, "error"))
		return
	}
	c.JSON(200, SuccessResponse(todo, "GET success!"))
}
func GetNumofTodo(c *gin.Context) {
	count, err := service.GetNumsofTodo()
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(0, "error"))
		return
	}
	c.JSON(200, SuccessResponse(count, "Count obtained successfully "))

}
