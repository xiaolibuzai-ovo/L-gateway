package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/consts"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
	"github.com/xiaolibuzai-ovo/L-gateway/gateWay/flow_calculate"
)

// HTTPFlowCalculateMiddleware 流量统计
func HTTPFlowCalculateMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		serverInterface, ok := c.Get("service")
		if !ok {
			// TODO 记录日志
			c.Abort()
			return
		}
		serviceDetail := serverInterface.(*dao.ServiceDetail)
		//统计项 1 全站 2 服务
		totalCounter, err := flow_calculate.FlowCalculateHandler.GetCounter(consts.FlowCalculateTotal)
		if err != nil {
			// TODO 记录日志
			c.Abort()
			return
		}
		totalCounter.Increase()

		//dayCount, _ := totalCounter.GetDayData(time.Now())
		//fmt.Printf("totalCounter qps:%v,dayCount:%v", totalCounter.QPS, dayCount)
		serviceCounter, err := flow_calculate.FlowCalculateHandler.GetCounter(consts.FlowCalculateServicePrefix + serviceDetail.Info.ServiceName)
		if err != nil {
			// TODO 记录日志
			c.Abort()
			return
		}
		serviceCounter.Increase()

		//dayServiceCount, _ := serviceCounter.GetDayData(time.Now())
		//fmt.Printf("serviceCounter qps:%v,dayCount:%v", serviceCounter.QPS, dayServiceCount)
		c.Next()
	}
}
