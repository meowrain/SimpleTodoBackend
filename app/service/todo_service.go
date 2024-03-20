package service

import (
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
	result := db.Delete(&models.Todo{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
} //更新数据//
func UpdateTodo(id int, updateTodo *models.Todo) error {
	db := utils.ConnectDB()
	if err := db.Save(&updateTodo).Error; err != nil {
		return nil
	}
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
	dp := utils.ConnectDB()
	var count int64
	if err := dp.Model(&models.Todo{}).Count(&count).Error; err != nil {
		return 0, err
	}
	return int(count), nil
}
