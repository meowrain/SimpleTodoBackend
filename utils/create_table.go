package utils

import (
	"fmt"
	"todoBackend/app/models"
)

func CreateTable() {
	// Database initialization
	db := ConnectDB()
	hasTableTodo := db.Migrator().HasTable(&models.Todo{})
	hasTableUser := db.Migrator().HasTable(&models.User{})
	if hasTableTodo && hasTableUser {
		fmt.Println("表已经存在")
		return
	}
	err := db.AutoMigrate(&models.User{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}
}
