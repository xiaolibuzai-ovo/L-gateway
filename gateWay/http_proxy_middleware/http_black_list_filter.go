package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
	"github.com/xiaolibuzai-ovo/L-gateway/utils"
	"strings"
)

// HTTPBlackListFilterMiddleware 匹配接入方式 基于请求信息
func HTTPBlackListFilterMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			// TODO 记录日志
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)

		var ipWhiteList []string
		if serviceDetail.AccessControl.WhiteList != "" {
			ipWhiteList = strings.Split(serviceDetail.AccessControl.WhiteList, ",")
		}
		// 未配置白名单才会校验黑名单
		currentIp := c.ClientIP()

		var ipBlackList []string
		if serviceDetail.AccessControl.BlackList != "" {
			ipBlackList = strings.Split(serviceDetail.AccessControl.BlackList, ",")
		}
		if serviceDetail.AccessControl.OpenAuth == 1 && len(ipWhiteList) == 0 && len(ipBlackList) > 0 {
			if utils.ContainsString(ipBlackList, currentIp) {
				// TODO 记录日志
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
