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

func CreateUser(u *models.User) error {
	var err error
	err = BeforeSave(u)
	//限制bio长度
	if err != nil {
		return err
	}
	if len(u.Bio) > 250 {
		return errors.New("Bio is too long, maximum is 240 characters")
	}
	db := utils.ConnectDB()
	err = db.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

// SaveUser
func SaveUser(u *models.User) error {
	var err error
	db := utils.ConnectDB()
	err = db.Save(&u).Error
	if err != nil {
		return err
	}
	return nil
}
func BeforeSave(u *models.User) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(u.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	return nil
}
func UpdateUser(inputUser, userFromDB *models.User) error {
	//这个可以直接SaveUser(u) 但是这样的话，会把所有的字段都更新，不太好
	if inputUser.Username != "" {
		userFromDB.Username = inputUser.Username
	}

	if inputUser.PasswordHash != "" {
		// 如果传进来的有值，就进行预处理
		err := BeforeSave(inputUser)
		if err != nil {
			return err
		}
		userFromDB.PasswordHash = inputUser.PasswordHash
	}
	if inputUser.Bio != "" {
		userFromDB.Bio = inputUser.Bio
	}

	if inputUser.Birthday != "" {
		userFromDB.Birthday = inputUser.Birthday
	}

	if inputUser.Email != "" {
		userFromDB.Email = inputUser.Email
	}
	if inputUser.Gender != "" {
		userFromDB.Gender = inputUser.Gender
	}

	if inputUser.PhoneNumber != "" {
		userFromDB.PhoneNumber = inputUser.PhoneNumber
	}
	// 更新时间戳
	userFromDB.UpdatedAt = time.Now()
	err := SaveUser(userFromDB)
	if err != nil {
		return err
	}
	return nil
}
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(u *models.User) (string, error) {
	var err error
	userInDB := models.User{} //加密过得
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

// GetUserByID
func GetUserByID(id uint) (models.User, error) {
	var u models.User
	db := utils.ConnectDB()
	if err := db.First(&u, id).Error; err != nil {
		return u, errors.New("user not found")
	}
	u.PasswordHash = ""
	return u, nil
}

func UpdateAvatar(u *models.User, avatarURL string) error {
	u.Avatar = avatarURL
	//更新用户信息
	err := SaveUser(u)
	if err != nil {
		return err
	}
	return nil
}

// 更新bio
func Updatebio(u *models.User, bio string) error {
	if len(bio) > 250 {
		return errors.New("Bio is too long")
	}
	u.Bio = bio
	err := SaveUser(u)
	if err != nil {
		return err
	}
	return nil
}

// 删除bio
func Dlebio(u *models.User) error {
	u.Bio = ""
	err := SaveUser(u)
	if err != nil {
		return err
	}
	return nil
}
