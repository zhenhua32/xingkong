package main

import (
	"net/http"
	"time"

	"github.com/zhenhua32/xingkong/pkg/logger"
	"github.com/zhenhua32/xingkong/web/app"
)

// @title 行空 API
// @version 1.0
// @description 行空 API, 用于小说搜索
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://xingkong.io/support
// @contact.email zhenhua32@xingkong.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /

func main() {
	var addr string = ":8080"

	srv := &http.Server{
		Addr:         addr,
		Handler:      app.Router(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	logger.Sugar.Infof("正在启动 Web 服务, 地址是 %s", addr)
	if err := srv.ListenAndServe(); err != nil {
		logger.Sugar.Panicf("启动 Web 服务失败, 原计划在 %s 上启动, 错误是 %v", addr, err)
	}

}
