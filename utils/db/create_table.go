package db

import (
	"todoBackend/app/api/todo/models"
	models2 "todoBackend/app/api/user/models"
	"todoBackend/utils/loggers"
)

// CreateTable 用于在数据库中创建表
func CreateTable() {
	// 数据库初始化
	hasTableTodo := DB.Migrator().HasTable(&models.Todo{})
	hasTableUser := DB.Migrator().HasTable(&models2.User{})
	if hasTableTodo && hasTableUser {
		return
	}
	err := DB.AutoMigrate(&models2.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&models.Todo{})
	if err != nil {
		return
	}

	loggers.TodoLogger.Info("表创建成功！")

}
