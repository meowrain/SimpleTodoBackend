package db

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"todoBackend/app/config"
	"todoBackend/utils/loggers"
)

var DB *gorm.DB

func init() {
	// 读取配置文件config.yaml
	host := config.Cfg.DB.Host
	port := config.Cfg.DB.Port
	user := config.Cfg.DB.User
	dbName := config.Cfg.DB.Name
	password := config.Cfg.DB.Password
	// 创建数据源名称
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, password, host, port, dbName)
	ConnectDB(dsn)
}

// ConnectDB 用于连接数据库并返回 *gorm.DB 实例
func ConnectDB(datasource string) *gorm.DB {
	// 创建新的日志记录器实例
	//newLogger := logger.New(
	//	log.New(os.Stdout, "\r\n", log.LstdFlags),
	//	logger.Config{
	//		SlowThreshold:             time.Second, // 设置慢查询阈值
	//		LogLevel:                  logger.Info, // 设置日志级别
	//		IgnoreRecordNotFoundError: true,        // 忽略未找到记录的错误
	//		ParameterizedQueries:      true,        // 启用参数化查询
	//		Colorful:                  false,       // 禁用彩色打印
	//	},
	//)

	var err error
	// 连接数据库
	DB, err = gorm.Open(mysql.Open(datasource), &gorm.Config{})
	if err != nil {
		panic("连接mysql数据库失败, error=" + err.Error())
	} else {
		loggers.Info("mysql连接成功")
	}
	sqlDB, err := DB.DB()
	if err != nil {
		panic("failed to get database connection pool")
	}
	sqlDB.SetMaxIdleConns(10)  // 设置最大空闲连接数
	sqlDB.SetMaxOpenConns(100) // 设置最大打开连接数
	return DB
}
