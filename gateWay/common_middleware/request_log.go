package middleware

import (
	"github.com/gin-gonic/gin"
)

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 打印请求参数
		c.Next()
	}
}
