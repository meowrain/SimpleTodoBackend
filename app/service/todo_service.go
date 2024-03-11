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
func DeleteTodo(todo *models.Todo) error {

	return nil
}
