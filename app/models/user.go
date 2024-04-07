package models

import (
	"gorm.io/gorm"
)

// User 模型代表用户表
type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`                  // 用户名，非空且唯一
	PasswordHash string `gorm:"not null" json:"password" binding:"-"`                                // 密码哈希，不进行JSON序列化和绑定
	Avatar       string `gorm:"default:'http://127.0.0.1:8090/users/static/user.svg'" json:"avatar"` // 头像链接，默认为用户.svg
	Bio          string `gorm:"type:text" json:"bio" validate:"max=250"`                             // 个人简介，最大长度250
	Email        string `json:"email"  gorm:"type:varchar(254)"`                                     // 电子邮件
	PhoneNumber  string `json:"phonenumber" validate:"max=20"`                                       // 电话号码，最大长度20
	Gender       string `json:"gender" gorm:"type:enum('男', '女');default:'男'"`                       // 性别，枚举值为'男'或'女'，默认为'男'
	Birthday     string `json:"birthday" gorm:"type:varchar(25)"`                                    // 生日
}
