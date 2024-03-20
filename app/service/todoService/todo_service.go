package todoService

import (
	"errors"
	"todoBackend/app/models"
	"todoBackend/utils"
)

// 创造todo
func CreateTodo(todo *models.Todo) error {
	db := utils.ConnectDB()
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
} //查看所有的todo
func GetAllTodo() ([]models.Todo, error) {
	db := utils.ConnectDB()
	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return todos, nil
	}
} //删除todo根据提供的ID
func DeleteTodo(id int) error {
	db := utils.ConnectDB()

	// 首先，尝试获取todo
	todo := models.Todo{}
	err := db.First(&todo, id).Error
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
	result := db.Delete(&todo)
	if err := result.Error; err != nil {
		// 如果在试图删除todo时有错误发生，返回这个错误
		return err
	}

	// 如果一切正常，返回nil表示没有错误
	return nil
}

// 更新数据//
func UpdateTodo(id int, updateTodo *models.Todo) error {
	db := utils.ConnectDB()
	//查询是否存在这个todo
	var todo models.Todo
	if err := db.First(&todo, id).Error; err != nil {
		return err
	}
	if err := db.Model(&todo).Updates(*updateTodo).Error; err != nil {
		return err
	}
	// 如果一个均无错误，则将更新的todo赋值给传入的结构，这样调用者可以获取到最新信息
	*updateTodo = todo
	return nil
}
func GetTodo(id int) (*models.Todo, error) {
	db := utils.ConnectDB()
	var todo models.Todo
	if err := db.First(&todo, id).Error; err != nil {
		return nil, err
	}
	return &todo, nil
}
func GetNumsofTodo() (int, error) {
	db := utils.ConnectDB()
	var count int64
	if err := db.Model(&models.Todo{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
