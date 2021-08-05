package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/errno"
	"github.com/zhenhua32/xingkong/pkg/logger"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

// JSON 返回 JSON 类型的 Response
func JSON(c *gin.Context, err error, data interface{}) {
	if err != nil {
		logger.Sugar.Errorf(err.Error())
	}
	code, msg := errno.DecodeErr(err)

	c.JSON(200, Response{Code: code, Msg: msg, Data: data})
}
