package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`
	PasswordHash string `gorm:"not null" json:"password" binding:"required"`
	Avatar       string `gorm:"default:'http://127.0.0.1:8090/users/static/user.svg'" json:"avatar"`
	Todos        []Todo `json:"todos"`
	Bio          string `grom:"typr:text" json:"bio" validate:"max=250"`
}
