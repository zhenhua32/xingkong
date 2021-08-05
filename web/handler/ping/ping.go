package ping

import "github.com/gin-gonic/gin"

type PingResp struct {
	Hello string `json:"hello"`
}

// Ping 测试连接是否正常
// @Summary 测试连接是否正常
// @Description 测试连接是否正常
// @ID ping
// @Tags 监控
// @Accept  html
// @Produce  json
// @Success 200 {object} ping.PingResp
// @Router /ping [get]
func Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"hello": "world",
	})
}
