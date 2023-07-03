package http_proxy_router

import (
	"log"
	"net/http"
	"time"
)

func HttpServerRun() {
	r := InitRouter()
	addr := ":8080"
	server := &http.Server{
		Addr:              addr,
		Handler:           r,
		ReadTimeout:       100 * time.Second, // TODO 所有超时可配置化
		ReadHeaderTimeout: 100 * time.Second,
		WriteTimeout:      100 * time.Second,
		IdleTimeout:       100 * time.Second,
		MaxHeaderBytes:    1 << 20,
	}
	log.Printf("[HttpServerRun] http_proxy_run %s", addr)
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("[HttpServerRun] http_proxy_run %s err:%v", addr, err)
	}
}
