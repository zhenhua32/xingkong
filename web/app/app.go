package app

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/logger"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
)

func RunSwaggerCmd() {
	command := "swag init --output api --generalInfo ./cmd/server/main.go"

	fmt.Println("准备生成 swagger doc, 执行的命令是", command)
	params := strings.Split(command, " ")
	cmd := exec.Command(params[0], params[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s\n", stdoutStderr)
}

func prepare() {
	RunSwaggerCmd()
	model.InitDB()
	model.SetUpDB()
}

func Router() http.Handler {
	prepare()

	r := gin.New()

	r.Use(gin.Logger())

	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	// 添加路由
	loadAPI(r)

	return r
}
