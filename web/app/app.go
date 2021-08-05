package app

import (
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/logger"
)

func Router() http.Handler {
	r := gin.New()

	r.Use(gin.Logger())

	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	// 添加路由
	loadAPI(r)

	return r
}
