package middleware

import (
	"github.com/gin-gonic/gin"
	"todoBackend/utils"
)

func ConnectDBMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := utils.ConnectDB()
		c.Set("db", db)
		c.Next()
	}
}
