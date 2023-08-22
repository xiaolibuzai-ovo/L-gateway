package http_proxy_middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/dao"
	"github.com/xiaolibuzai-ovo/L-gateway/gateWay/reverse_proxy"
)

// HTTPReverseProxyMiddleware 反向代理 匹配接入方式 基于请求信息
func HTTPReverseProxyMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		service, ok := c.Get("service")
		if !ok {
			// TODO 记录日志
			c.Abort()
			return
		}
		serviceDetail := service.(*dao.ServiceDetail)

		//负载均衡
		lb, err := dao.LoadBalancerManipulator.GetLoadBalancer(serviceDetail)
		if err != nil {
			// TODO 记录日志
			c.Abort()
			return
		}

		//返回http连接池
		trans, err := dao.TransportManipulator.GetTrans(serviceDetail)
		if err != nil {
			// TODO 记录日志
			c.Abort()
			return
		}
		//创建 reverse_proxy
		reverseProxy := reverse_proxy.NewLoadBalanceReverseProxy(c, lb, trans)
		//使用 reverse_proxy.ServerHTTP(c.Request,c.Response)
		reverseProxy.ServeHTTP(c.Writer, c.Request)
		c.Abort()
		return

	}
}
