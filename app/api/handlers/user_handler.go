package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/app/models"
	"todoBackend/app/service"
	"todoBackend/utils"
	"todoBackend/utils/token"
)

// Register 用于处理用户注册请求。
func Register(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), "error"))
		return
	}
	err := service.CreateUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), "error"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(inputUser, "success"))
}

// Login 用于处理用户登录的请求
func Login(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), "error"))
		return
	}
	loginCheck, err := service.LoginCheck(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), "登录失败"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(loginCheck, "get loginCheck success!"))
}

// CurrentUser 用来获取当前用户的信息并返回给前端
func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := service.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(u, "success"))
}

// UpdateUser 用来更新用户信息
func UpdateUser(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, utils.ErrorResponse(err.Error(), "error"))
		return
	}
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFromDB, err := service.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, utils.NotFoundResponse(err.Error()))
		return
	}
	if err := service.UpdateUser(&inputUser, &userFromDB); err != nil {
		c.JSON(http.StatusNotFound, utils.ErrorResponse(err.Error(), "update failed"))
		return
	}
	c.JSON(http.StatusOK, utils.SuccessResponse(userFromDB, "success update"))

}
