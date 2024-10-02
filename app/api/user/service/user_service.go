package service

import (
	"errors"
	"html"
	"strings"
	"time"
	"todoBackend/app/api/user/models"
	"todoBackend/utils/db"
	"todoBackend/utils/jwts"
	"todoBackend/utils/pwd"
)

// CreateUser 创建新用户
func CreateUser(u *models.User) error {
	var err error
	hashPassword := pwd.HashPassword(u.Password)
	u.Password = hashPassword
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	// 限制bio长度
	if len(u.Bio) > 250 {
		return errors.New("个人简介太长了，最多只能有240个字符。")
	}
	err = db.DB.Create(&u).Error
	if err != nil {
		return errors.New("数据库创建用户失败")
	}
	return nil
}

// UpdateUser 更新用户信息
func UpdateUser(inputUser, userFromDB *models.User) error {
	updatesMap := map[string]interface{}{
		"Username":    inputUser.Username,
		"Bio":         inputUser.Bio,
		"Email":       inputUser.Email,
		"PhoneNumber": inputUser.PhoneNumber,
		"Gender":      inputUser.Gender,
		"Birthday":    inputUser.Birthday,
		"UpdatedAt":   time.Now(),
	}
	err := db.DB.Model(&userFromDB).Updates(updatesMap).Error
	if err != nil {
		return err
	}
	return nil
}

// LoginCheck 登录验证
func LoginCheck(u *models.User) (string, error) {
	var err error
	userInDB := models.User{} // 数据库中存储的user
	err = db.DB.Model(models.User{}).Where("username = ?", u.Username).Take(&userInDB).Error
	if err != nil {
		return "", errors.New("此用户不存在，请先注册！")
	}
	var jwtToken string
	if pwd.CheckPasswordHash(u.Password, userInDB.Password) {
		//如果密码校验通过
		jwtToken, err = jwts.GenerateToken(jwts.JwtPayload{
			UserId:   int(userInDB.ID),
			NickName: userInDB.Username,
		})
		if err != nil {
			return "", err
		}
		return jwtToken, nil
	}

	return "", errors.New("jwt token生成失败222")
}

// GetUserByID 通过ID获取用户信息
func GetUserByID(id uint) (models.User, error) {
	var u models.User
	if err := db.DB.First(&u, id).Error; err != nil {
		return u, errors.New("user_model not found")
	}
	u.Password = ""
	return u, nil
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(u *models.User, avatarURL string) error {
	// 更新用户信息
	err := db.DB.Model(&u).Update("avatar", avatarURL).Error
	if err != nil {
		return err
	}
	return nil
}
