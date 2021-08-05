package app

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/web/handler/ping"
)

func loadAPI(g *gin.Engine) {
	g.GET("/ping", ping.Ping)
}
