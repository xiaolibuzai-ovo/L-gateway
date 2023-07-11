package http_proxy_router

import (
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/GateWay/http_proxy_middleware"
	"net/http"
)

func InitRouter(middlewares ...gin.HandlerFunc) *gin.Engine {
	router := gin.New()
	router.Use(middlewares...)
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	router.Use(
		http_proxy_middleware.HTTPAccessWayMiddleware(),       // 匹配接入方式
		http_proxy_middleware.HTTPFlowCalculateMiddleware(),   // 流量统计
		http_proxy_middleware.HTTPFlowLimitMiddleware(),       // 限流
		http_proxy_middleware.HTTPWhiteListFilterMiddleware(), // 白名单过滤
		http_proxy_middleware.HTTPBlackListFilterMiddleware(), // 黑名单过滤
		http_proxy_middleware.HTTPReverseProxyMiddleware(),    // 反向代理
	)
	return router
}
