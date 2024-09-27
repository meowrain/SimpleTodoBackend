package todo_model

import "gorm.io/gorm"

// Todo 结构表示了一个任务
type Todo struct {
	gorm.Model        // gorm 提供的基本 model 结构
	Content    string `json:"content"` // 任务内容
	Status     int    `json:"status"`  // 任务状态
	UserID     uint   `json:"userID"`  // 用户ID
	Tag        string `json:"tag"`     // 任务标签
}
