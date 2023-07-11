package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/GateWay/flow_limit"
	"github.com/xiaolibuzai-ovo/L-gateway/consts"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
)

func HTTPFlowLimitMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			// TODO 记录日志
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		if serviceDetail.AccessControl.ServiceFlowLimit != 0 {
			serviceLimiter, err := flow_limit.FlowLimiterHandler.GetLimiter(
				consts.FlowCalculateServicePrefix+serviceDetail.Info.ServiceName,
				float64(serviceDetail.AccessControl.ServiceFlowLimit))
			if err != nil {
				// TODO 记录日志
				c.Abort()
				return
			}
			if !serviceLimiter.Allow() {
				// TODO 记录日志
				c.Abort()
				return
			}
		}

		if serviceDetail.AccessControl.ClientIPFlowLimit > 0 {
			clientLimiter, err := flow_limit.FlowLimiterHandler.GetLimiter(
				consts.FlowServiceLimitPrefix+serviceDetail.Info.ServiceName+"_"+c.ClientIP(),
				float64(serviceDetail.AccessControl.ClientIPFlowLimit))
			if err != nil {
				// TODO 记录日志
				c.Abort()
				return
			}
			if !clientLimiter.Allow() {
				// TODO 记录日志
				c.Abort()
				return
			}
		}
		c.Next()
	}
}
