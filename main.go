package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/xiaolibuzai-ovo/L-gateway/GateWay/http_proxy_router"
	dashboard "github.com/xiaolibuzai-ovo/L-gateway/biz"
	"github.com/xiaolibuzai-ovo/L-gateway/biz/dal"
	"github.com/xiaolibuzai-ovo/L-gateway/biz/dao"
	"os"
)

// config
const (
	port = 8080
	size = 20 << 20
)

//endpoint dashboard 后台管理 server gateway

var (
	endpoint = flag.String("endpoint", "", "input endpoint dashboard or server")
)

func main() {
	flag.Parse()
	if *endpoint == "" {
		flag.Usage()
		os.Exit(1)
	}
	dal.Init()

	if *endpoint == "dashboard" {
		//初始化数据库
		r := gin.Default()
		//h := server.Default(
		//	//server.WithHostPorts(port),
		//	server.WithMaxRequestBodySize(size),
		//	server.WithTransport(standard.NewTransporter),
		//)
		dashboard.RegisterDashBoard(r)

		r.Run(fmt.Sprintf(":%d", port)) // 监听并在 0.0.0.0:8080 上启动服务
	} else if *endpoint == "server" {
		// 加载服务列表到内存中,优化速度
		ctx := context.Background()
		dao.ServiceManipulator.LoadServiceManager(ctx)
		go func() {
			//启动http端口监听
			http_proxy_router.HttpServerRun()
		}()
		go func() {
			//启动https端口监听
		}()

	}
}
