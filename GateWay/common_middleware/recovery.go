package middleware

import (
	"github.com/gin-gonic/gin"
)

// RecoveryMiddleware 捕获所有panic，并且返回错误信息
func RecoveryMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// 日志记录
			}
		}()
		c.Next()
	}
}
