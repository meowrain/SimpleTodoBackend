package service

import (
	"todoBackend/app/api/todo/models"
	"todoBackend/utils/db"
)

// GetAllTodo 获取所有属于指定用户的todo列表
func GetAllTodo(userId uint) ([]models.Todo, error) {
	var todos []models.Todo
	if err := db.DB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

func AddAllTodo(todoList []models.TodoRequest, userId uint) error {
	tx := db.DB.Begin() // 开启事务
	if tx.Error != nil {
		return tx.Error
	}
	// 删除当前用户的所有数据
	if err := tx.Where("user_id = ?", userId).Delete(&models.Todo{}).Error; err != nil {
		tx.Rollback() // 回滚事务
		return err
	}

	for _, v := range todoList {
		var todo models.Todo = models.Todo{
			Content:   v.Content,
			Tag:       v.Tag,
			Completed: v.Completed,
			UserID:    userId,
		}
		if err := tx.Create(&todo).Error; err != nil {
			tx.Rollback() // 回滚事务
			return err
		}
	}
	if err := tx.Commit().Error; err != nil {
		return err
	}
	return nil
}
