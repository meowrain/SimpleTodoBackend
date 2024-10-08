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
func UpdateTodoList(todoList []models.TodoRequest, userId uint) error {
	tx := db.DB.Begin() // 开启事务
	if tx.Error != nil {
		return tx.Error
	}
	//首先去数据库找到现在所有的todo
	var currentTodosInDB []models.Todo
	if err := tx.Where("user_id = ?", userId).Find(&currentTodosInDB).Error; err != nil {
		tx.Rollback()
		return err
	}
	//接下来构建一个新的map，content 对应一个models.Todo实例
	currentTodoMap := make(map[string]models.Todo)
	for _, todo := range currentTodosInDB {
		currentTodoMap[todo.Content] = todo
	}

	// 我们在controller层获取了TodoRequestList，我们遍历这个List创建一个map,一个content对应一个models.TodoRequest实例
	newTodoMap := make(map[string]models.TodoRequest)
	for _, todo := range todoList {
		newTodoMap[todo.Content] = todo
	}
	// 遍历上面构建的传进来的待更新数据的map
	for task, newTodo := range newTodoMap {
		//如果在数据库种这个task存在
		if currentTodo, exists := currentTodoMap[task]; exists {
			//就把传进来的待更新数据的Completed字段和Tag字段更新，然后使用事务保存
			currentTodo.Completed = newTodo.Completed
			currentTodo.Tag = newTodo.Tag
			if err := tx.Save(&currentTodo).Error; err != nil {
				tx.Rollback()
				return err
			}

		} else {
			// 如果这个task不存在，那么我们就执行添加操作
			// 调取newTodo中的数据构建models.Todo实例
			newTodoModel := models.Todo{
				UserID:    userId,
				Content:   newTodo.Content,
				Completed: newTodo.Completed,
				Tag:       newTodo.Tag,
			}
			// 使用事务在数据库种创建相关数据
			if err := tx.Create(&newTodoModel).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	//如果传入的新数据和数据库中的数据比对后发现有个task不见了，那就从数据库中删除这行记录
	for task, currentTodo := range currentTodoMap {
		if _, exists := newTodoMap[task]; !exists {
			if err := tx.Delete(&currentTodo).Error; err != nil {
				tx.Rollback()
				return err
			}
		}
	}
	//如果上面的事务都没有遇到问题，那么就提交事务，把数据更新到数据库中
	if err := tx.Commit().Error; err != nil {
		return err
	}

	return nil
}
