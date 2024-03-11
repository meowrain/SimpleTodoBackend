package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todoBackend/app/config"
)

func ConnectDB() *gorm.DB {
	//读取配置文件config.yaml
	host := config.Cfg.DB.Host
	port := config.Cfg.DB.Port
	user := config.Cfg.DB.User
	dbName := config.Cfg.DB.Name
	password := config.Cfg.DB.Password
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
