package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`
	PasswordHash string `gorm:"not null" json:"password" binding:"-"`
	Avatar       string `gorm:"default:'http://127.0.0.1:8090/users/static/user.svg'" json:"avatar"`
	Bio          string `gorm:"type:text" json:"bio" validate:"max=250"`
	Email        string `json:"email"  gorm:"type:varchar(254)"`
	PhoneNumber  string `json:"phonenumber" validate:"max=20"`
	Gender       string `json:"gender" gorm:"type:enum('男', '女');default:'男'"`
	Birthday     string `json:"birthday" gorm:"type:varchar(25)"`
}
