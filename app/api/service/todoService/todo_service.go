package todoService

import (
	"errors"
	"todoBackend/app/models/todo_model"
	"todoBackend/utils/db"
)

// CreateTodo 创建一个新的todo
func CreateTodo(todo *todo_model.Todo) error {
	if err := db.DB.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}

// GetAllTodo 获取所有属于指定用户的todo列表
func GetAllTodo(userId uint) ([]todo_model.Todo, error) {
	var todos []todo_model.Todo
	if err := db.DB.Where("user_id = ?", userId).Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}

// DeleteTodo 根据提供的ID删除todo
func DeleteTodo(id int) error {

	// 首先，尝试获取todo
	todo := todo_model.Todo{}
	err := db.DB.First(&todo, id).Error
	if err != nil {
		// 如果我们无法找到这个todo，报错
		return errors.New("Todo未找到")
	}

	// 然后检查todo是否已经被删除
	if todo.DeletedAt.Valid {
		// 如果我们找到的todo已经被删除，报错
		return errors.New("这个todo已经被删除过了")
	}

	// 如果todo存在并且未被删除，我们就删除它
	result := db.DB.Delete(&todo)
	if err := result.Error; err != nil {
		// 如果在试图删除todo时有错误发生，返回这个错误
		return err
	}

	// 如果一切正常，返回nil表示没有错误
	return nil
}

// UpdateTodo 根据ID更新todo数据
func UpdateTodo(id int, updateTodo *todo_model.Todo) error {
	//查询是否存在这个todo
	var todo todo_model.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		return err
	}
	if err := db.DB.Model(&todo).Updates(*updateTodo).Error; err != nil {
		return err
	}
	// 如果一个均无错误，则将更新的todo赋值给传入的结构，这样调用者可以获取到最新信息
	*updateTodo = todo
	return nil
}

// GetTodo 根据ID获取特定的todo
func GetTodo(id int) (*todo_model.Todo, error) {
	var todo todo_model.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}

// GetNumsofTodo 获取特定用户的todo数量
func GetNumsofTodo(userId uint) (int, error) {
	var count int64
	if err := db.DB.Model(&todo_model.Todo{}).Where("user_id = ?", userId).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
