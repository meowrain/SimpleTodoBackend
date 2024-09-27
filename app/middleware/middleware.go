package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/utils/jwts"
)

// JwtAuthMiddleware JwtAuthMiddleware是一个JWT认证中间件
// 它调用了token包中的Valid()函数以校验请求中的JWT jwts。
// 如果校验失败，它就会返回一个401 Unauthorized状态码，并且中断后续的中间件函数或者处理函数。
// 如果token校验通过，它就会调用c.Next()函数以便后续的中间件函数和处理器可以继续执行。
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := jwts.Valid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized") //如果校验失败，则返回401状态码
			c.Abort()                                         //中断后续的函数
			return
		}
		c.Next() //如果校验通过，调用c.Next()函数
	}
}
