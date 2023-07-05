package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
)

// HTTPAccessWayMiddleware 匹配接入方式 基于请求信息
func HTTPAccessWayMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取接入配置
		service, err := dao.ServiceManipulator.HTTPAccessWay(c)
		if err != nil {
			// TODO 记录日志
			c.Abort()
			return
		}
		c.Set("service", service)
		c.Next()
	}
}
