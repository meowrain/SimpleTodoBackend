package db

import (
	"todoBackend/app/models/feedback_model"
	"todoBackend/app/models/todo_model"
	"todoBackend/app/models/user_model"
	"todoBackend/utils/loggers"
)

// CreateTable 用于在数据库中创建表
func CreateTable() {
	// 数据库初始化
	hasTableTodo := DB.Migrator().HasTable(&todo_model.Todo{})
	hasTableUser := DB.Migrator().HasTable(&user_model.User{})
	hasTableFeedBack := DB.Migrator().HasTable(&feedback_model.FeedBack{})
	hasComments := DB.Migrator().HasTable(&feedback_model.Comment{})
	if hasTableTodo && hasTableUser && hasTableFeedBack && hasComments {
		return
	}
	err := DB.AutoMigrate(&user_model.User{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&todo_model.Todo{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&feedback_model.FeedBack{})
	if err != nil {
		return
	}
	err = DB.AutoMigrate(&feedback_model.Comment{})
	if err != nil {
		return
	}
	loggers.TodoLogger.Info("表创建成功！")

}
