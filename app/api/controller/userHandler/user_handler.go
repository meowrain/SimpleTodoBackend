package userHandler

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"todoBackend/app/api/service/userService"
	"todoBackend/app/config"
	"todoBackend/app/models"
	"todoBackend/utils/responses"
	"todoBackend/utils/token"

	"github.com/gin-gonic/gin"
)

// Register 用于处理用户注册请求。
func Register(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}
	err := userService.CreateUser(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse(inputUser, "success"))
}

// Login 用于处理用户登录的请求
func Login(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}
	loginCheck, err := userService.LoginCheck(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "登录失败"))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse(loginCheck, "get loginCheck success!"))
}

// CurrentUser 用来获取当前用户的信息并返回给前端
func CurrentUser(c *gin.Context) {
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := userService.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse(u, "success"))
}

// UpdateUser 用来更新用户信息
func UpdateUser(c *gin.Context) {
	var inputUser models.User
	if err := c.ShouldBindJSON(&inputUser); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}
	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFromDB, err := userService.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NotFoundResponse(err.Error()))
		return
	}
	if err := userService.UpdateUser(&inputUser, &userFromDB); err != nil {
		c.JSON(http.StatusNotFound, responses.ErrorResponse(err.Error(), "update failed"))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse(userFromDB, "success update"))

}

// UploadAvatar 用于处理上传头像的请求
func UploadAvatar(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, must-revalidate")
	c.Header("Pragma", "no-cache")
	c.Header("Expires", "0")

	userId, err := token.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	file, err := c.FormFile("avatar")
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}

	fileName := file.Filename
	extName := filepath.Ext(fileName)
	userFromDB, err := userService.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}

	err = c.SaveUploadedFile(file, "app/static/avatars/"+strconv.Itoa(int(userId))+extName)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error(), "保存文件失败"))
		return
	}

	server := config.Cfg.Server.URL
	uploadUrl := fmt.Sprintf("%s/users/avatars/", server) + strconv.Itoa(int(userId)) + extName
	err = userService.UpdateAvatar(&userFromDB, uploadUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error(), "更新用户头像失败"))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse(userFromDB, "avatar uploaded successfully"))
}
