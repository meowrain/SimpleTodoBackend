package service

import (
	"todoBackend/app/models"
	"todoBackend/utils"
)

func CreateTodo(todo *models.Todo) error {
	db := utils.ConnectDB()
	if err := db.Create(&todo).Error; err != nil {
		return err
	}
	return nil
}
func GetAllTodo() ([]models.Todo, error) {
	db := utils.ConnectDB()
	var todos []models.Todo
	if err := db.Find(&todos).Error; err != nil {
		return nil, err
	} else {
		return todos, nil
	}
}
func DeleteTodo(id int) error {
	db := utils.ConnectDB()
	result := db.Delete(&models.Todo{}, id)
	if err := result.Error; err != nil {
		return err
	}
	return nil
}
