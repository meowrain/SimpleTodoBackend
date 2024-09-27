package loggers

import (
	"github.com/sirupsen/logrus"
	"os"
)

var TodoLogger *logrus.Logger

func init() {
	TodoLogger = logrus.New()
	TodoLogger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp:          true,
		TimestampFormat:        "2006-01-02 15:04:05", // 自定义时间格式\
		DisableLevelTruncation: true,                  // 防止日志级别被截断
		ForceColors:            true,
	})
	TodoLogger.SetOutput(os.Stdout)
	TodoLogger.SetLevel(logrus.DebugLevel)
}

// 对外暴露日志记录函数
func Info(args ...interface{}) {
	TodoLogger.Info(args...)
}

func Warn(args ...interface{}) {
	TodoLogger.Warn(args...)
}

func Error(args ...interface{}) {
	TodoLogger.Error(args...)
}

func Debug(args ...interface{}) {
	TodoLogger.Debug(args...)
}
