package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/utils"
	"todoBackend/utils/token"
)

// ConnectDBMiddleWare ConnectDBMiddleWare是一个中间件函数，用于连接数据库
// 这是通过调用utils包中的ConnectDB()函数实现的。
// 它返回一个gin.HandlerFunc类型的函数，
// 在这个函数中，我们通过调用c.Set()函数将数据库连接存储在gin的上下文中。
// 这样，后续的处理器就能够通过从上下文中获取到数据库连接。
func ConnectDBMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := utils.ConnectDB()
		c.Set("db", db)
		c.Next()
	}
}

// JwtAuthMiddleware JwtAuthMiddleware是一个JWT认证中间件
// 它调用了token包中的Valid()函数以校验请求中的JWT token。
// 如果校验失败，它就会返回一个401 Unauthorized状态码，并且中断后续的中间件函数或者处理函数。
// 如果token校验通过，它就会调用c.Next()函数以便后续的中间件函数和处理器可以继续执行。
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.Valid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized") //如果校验失败，则返回401状态码
			c.Abort()                                         //中断后续的函数
			return
		}
		c.Next() //如果校验通过，调用c.Next()函数
	}
}
