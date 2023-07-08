package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
	"github.com/xiaolibuzai-ovo/L-gateway/utils"
	"strings"
)

// HTTPWhiteListFilterMiddleware ip白名单校验
func HTTPWhiteListFilterMiddleware() gin.HandlerFunc {
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
		if serviceDetail.AccessControl.OpenAuth == 1 && len(ipWhiteList) > 0 {
			if !utils.ContainsString(ipWhiteList, c.ClientIP()) {
				// TODO 记录日志
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
