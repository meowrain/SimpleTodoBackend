package controller

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"todoBackend/app/api/auth/auth_model"
	"todoBackend/app/api/user/models"
	"todoBackend/app/api/user/service"
	"todoBackend/app/config"
	"todoBackend/utils/jwts"
	"todoBackend/utils/responses"

	"github.com/gin-gonic/gin"
)

// Register 用于处理用户注册请求。
func Register(c *gin.Context) {
	var registerRequest auth_model.RegisterRequest
	if err := c.ShouldBindJSON(&registerRequest); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
	}
	var userToSave models.User = models.User{
		Username: registerRequest.Username,
		Password: registerRequest.PasswordHash,
	}
	err := service.CreateUser(&userToSave)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse(userToSave, "success"))
}

// Login 用于处理用户登录的请求
func Login(c *gin.Context) {
	var loginRequest auth_model.LoginRequest
	if err := c.ShouldBindJSON(&loginRequest); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "error"))
	}
	var inputUser models.User = models.User{
		Username: loginRequest.Username,
		Password: loginRequest.PasswordHash,
	}

	jwtToken, err := service.LoginCheck(&inputUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse(err.Error(), "登录失败"))
		return
	}
	c.Header("Authorization", "Bearer "+jwtToken)
	c.JSON(http.StatusOK, responses.SuccessResponse(jwtToken, "登录成功！"))
}

// CurrentUser 用来获取当前用户的信息并返回给前端
func CurrentUser(c *gin.Context) {
	userId, err := jwts.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u, err := service.GetUserByID(userId)
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
	userId, err := jwts.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userFromDB, err := service.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusNotFound, responses.NotFoundResponse(err.Error()))
		return
	}
	if err := service.UpdateUser(&inputUser, &userFromDB); err != nil {
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

	userId, err := jwts.ExtractTokenID(c)
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
	userFromDB, err := service.GetUserByID(userId)
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
	err = service.UpdateAvatar(&userFromDB, uploadUrl)
	if err != nil {
		c.JSON(http.StatusInternalServerError, responses.ErrorResponse(err.Error(), "更新用户头像失败"))
		return
	}

	c.JSON(http.StatusOK, responses.SuccessResponse(userFromDB, "avatar uploaded successfully"))
}

func ChanagePassword(c *gin.Context) {
	passwordRequest := auth_model.PasswordRequest{}
	// 从请求中解析JSON数据到passwordRequest结构体
	if err := c.ShouldBindJSON(&passwordRequest); err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse("", "Bad Request"))
		return
	}
	userId, err := jwts.ExtractTokenID(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse("", "extract token from jwt failed"))
		return
	}
	userFromDB, err := service.GetUserByID(userId)
	if err != nil {
		c.JSON(http.StatusBadRequest, responses.ErrorResponse("", "Can't get user info from database"))
		return
	}

	err = service.UpdateUserPassword(&userFromDB, passwordRequest.PasswordHash)
	if err != nil {
		c.JSON(http.StatusBadGateway, responses.ErrorResponse("", "update user password failed"))
		return
	}
	c.JSON(http.StatusOK, responses.SuccessResponse("", "Update User Password Successfully!"))

}
