package todoHandler

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
	"strconv"
	"todoBackend/app/api/service/todoService"
	"todoBackend/app/models/todo_model"
	"todoBackend/utils/jwts"
	. "todoBackend/utils/responses"
)

// CreateTodo 创建todo
func CreateTodo(c *gin.Context) {
	var todo todo_model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	if err := todoService.CreateTodo(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todo, "Add success!"))
}

// GetAllTodo 获取所有todo
func GetAllTodo(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	todos, err := todoService.GetAllTodo(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "获取全部todo失败"))
	}
	c.JSON(http.StatusOK, SuccessResponse(todos, "get successfully!"))
}

// DeleteTodo 删除todo
func DeleteTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := todoService.DeleteTodo(id); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "删除todo失败"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(id, "todo 删除成功"))
}

// UpdateTodo 更新todo
func UpdateTodo(c *gin.Context) {
	var todo todo_model.Todo
	if err := c.ShouldBindJSON(&todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := todoService.UpdateTodo(id, &todo); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err, "error"))
		return
	}
	c.JSON(http.StatusOK, SuccessResponse(todo, "Update success!"))
}

// GetTodo 获取单个todo
func GetTodo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	todo, err := todoService.GetTodo(id)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse(err, "error"))
		return
	}
	c.JSON(200, SuccessResponse(todo, "GET success!"))
}

// GetNumofTodo 获取todo数量
func GetNumofTodo(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	count, err := todoService.GetNumsofTodo(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(0, "error"))
		return
	}
	c.JSON(200, SuccessResponse(count, "Count obtained successfully "))
}

// UploadTodoPhoto 上传todo照片
func UploadTodoPhoto(c *gin.Context) {
	todoid, _ := strconv.Atoi(c.Param("id"))
	file, err := c.FormFile("photo")
	if err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse(err.Error(), "error"))
		return
	}
	//获取文件名的后缀
	extname := filepath.Ext(file.Filename)
	//将文件保存到指定的路径
	err = c.SaveUploadedFile(file, fmt.Sprintf("app/static/photots/%d%s", todoid, extname))
	if err != nil {
		c.JSON(500, "文件保存失败")
		return
	}
	c.JSON(200, SuccessResponse("上传成功", "Upload successfully"))
}
