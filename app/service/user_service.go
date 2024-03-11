package service

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"html"
	"strings"
	"todoBackend/app/models"
	"todoBackend/utils"
)

func SaveUser(u *models.User) error {
	var err error
	err = BeforeSave(u)
	if err != nil {
		return err
	}
	db := utils.ConnectDB()
	err = db.Create(&u).Error
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
	token, err := utils.GenerateToken(u.ID)
	if err != nil {
		return "", err
	}
	return token, nil

}
