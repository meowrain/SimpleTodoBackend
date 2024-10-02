package models

import (
	"gorm.io/gorm"
	"todoBackend/app/api/user/models"
)

// Todo 结构表示了一个任务
type Todo struct {
	gorm.Model             // gorm 提供的基本 model 结构
	Content    string      `json:"content"`                         // 任务内容
	Completed  bool        `json:"completed"`                       // 任务状态
	UserID     uint        `json:"user_id"`                         // 用户ID
	Tag        string      `json:"tag"`                             // 任务标签
	User       models.User `gorm:"foreignKey:UserID;references:ID"` // 外键关联
}
