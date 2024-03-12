package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"todoBackend/utils"
	"todoBackend/utils/token"
)

func ConnectDBMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		db := utils.ConnectDB()
		c.Set("db", db)
		c.Next()
	}
}
func JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := token.Valid(c)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		c.Next()
	}
}
