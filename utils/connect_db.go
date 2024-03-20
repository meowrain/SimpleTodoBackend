package utils

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
	"todoBackend/app/config"
)

var db *gorm.DB

func ConnectDB() *gorm.DB {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,        // Don't include params in the SQL log
			Colorful:                  false,       // Disable color
		},
	)
	if db == nil {
		//读取配置文件config.yaml
		host := config.Cfg.DB.Host
		port := config.Cfg.DB.Port
		user := config.Cfg.DB.User
		dbName := config.Cfg.DB.Name
		password := config.Cfg.DB.Password
		dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
		var err error
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
			Logger: newLogger,
		})
		if err != nil {
			panic("failed to connect database")
		}
		sqlDB, err := db.DB()
		if err != nil {
			panic("failed to get database connection pool")
		}
		sqlDB.SetMaxIdleConns(10)  // 设置最大空闲连接数
		sqlDB.SetMaxOpenConns(100) // 设置最大打开连接数
	}
	return db
}
