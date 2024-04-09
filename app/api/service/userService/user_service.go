package userService

import (
	"errors"
	"html"
	"strings"
	"time"
	"todoBackend/app/models"
	"todoBackend/utils"
	"todoBackend/utils/token"

	"golang.org/x/crypto/bcrypt"
)

// CreateUser 创建新用户
func CreateUser(u *models.User) error {
	var err error
	err = BeforeSave(u)
	// 限制bio长度
	if err != nil {
		return err
	}
	if len(u.Bio) > 250 {
		return errors.New("bio is too long, maximum is 240 characters")
	}
	db := utils.ConnectDB()
	err = db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveUser 保存用户信息
func SaveUser(u *models.User) error {
	var err error
	db := utils.ConnectDB()
	err = db.Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// BeforeSave 在保存用户前进行处理，如密码哈希化和去除用户名中的特殊字符
func BeforeSave(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
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
	db := utils.ConnectDB()
	err := db.Model(&userFromDB).Updates(updatesMap).Error
	if err != nil {
		return err
	}
	return nil
}

// VerifyPassword 验证密码是否匹配
func VerifyPassword(password, hashedPassword string) error {
	if hashedPassword == "" {
		// 如果hashedPassword是空的，返回一个错误
		return errors.New("hashed password is empty, cannot verify password")
	}
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// LoginCheck 登录验证
func LoginCheck(u *models.User) (string, error) {
	var err error
	userInDB := models.User{} // 加密过得
	db := utils.ConnectDB()
	err = db.Model(models.User{}).Where("username = ?", u.Username).Take(&userInDB).Error
	if err != nil {
		return "", err
	}
	err = VerifyPassword(u.PasswordHash, userInDB.PasswordHash)
	if err != nil && errors.Is(err, bcrypt.ErrMismatchedHashAndPassword) {
		return "", err
	}
	generateToken, err := token.GenerateToken(userInDB.ID)
	if err != nil {
		return "", err
	}
	return generateToken, nil
}

// GetUserByID 通过ID获取用户信息
func GetUserByID(id uint) (models.User, error) {
	var u models.User
	db := utils.ConnectDB()
	if err := db.First(&u, id).Error; err != nil {
		return u, errors.New("user not found")
	}
	u.PasswordHash = ""
	return u, nil
}

// UpdateAvatar 更新用户头像
func UpdateAvatar(u *models.User, avatarURL string) error {
	// 更新用户信息
	db := utils.ConnectDB()
	err := db.Model(&u).Update("avatar", avatarURL).Error
	if err != nil {
		return err
	}
	return nil
}
