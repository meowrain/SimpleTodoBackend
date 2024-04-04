package utils

import (
	"fmt"
	"todoBackend/app/models"
)

func CreateTable() {
	// 数据库初始化
	db := ConnectDB()
	hasTableTodo := db.Migrator().HasTable(&models.Todo{})
	hasTableUser := db.Migrator().HasTable(&models.User{})
	hasTableFeedBack := db.Migrator().HasTable(&models.FeedBack{})
	hasComments := db.Migrator().HasTable(&models.Comment{})
	if hasTableTodo && hasTableUser && hasTableFeedBack && hasComments {
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
	err = db.AutoMigrate(&models.FeedBack{})
	if err != nil {
		return
	}
	err = db.AutoMigrate(&models.Comment{})
	if err != nil {
		return
	}
}
