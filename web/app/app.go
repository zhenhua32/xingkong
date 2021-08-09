package app

import (
	"fmt"
	"net/http"
	"os/exec"
	"strings"
	"time"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/zhenhua32/xingkong/pkg/logger"
	model "github.com/zhenhua32/xingkong/pkg/model/gorm"
)

func RunSwaggerCmd() {
	// --parseDependency --parseInternal 加这两个参数后极慢 https://github.com/swaggo/swag/issues/810
	command := "swag init --output api --generalInfo ./cmd/server/main.go"

	fmt.Println("准备生成 swagger doc, 执行的命令是", command)
	params := strings.Split(command, " ")
	cmd := exec.Command(params[0], params[1:]...)
	stdoutStderr, err := cmd.CombinedOutput()
	fmt.Printf("%s\n", stdoutStderr)
	if err != nil {
		panic(err)
	}
}

func prepare() {
	// RunSwaggerCmd()
	model.InitDB()
	model.SetUpDB()
}

func loadCors(g *gin.Engine) {
	g.Use(cors.New(cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowOrigins:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
}

func Router() http.Handler {
	prepare()

	r := gin.New()

	// 日志配置
	r.Use(gin.Logger())

	r.Use(ginzap.Ginzap(logger.Logger, time.RFC3339, false))
	r.Use(ginzap.RecoveryWithZap(logger.Logger, true))

	// 配置 cors
	loadCors(r)

	// 添加路由
	loadAPI(r)

	return r
}
