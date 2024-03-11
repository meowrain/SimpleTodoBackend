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
	// Migrate the schema
	db.AutoMigrate(&models.User{})
	db.AutoMigrate(&models.Todo{})
}
