package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username     string `gorm:"not null;unique" json:"username" binding:"required"`
	PasswordHash string `gorm:"not null" json:"password" binding:"required"`
	Todos        []Todo `json:"todos"`
}
