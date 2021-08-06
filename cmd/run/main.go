package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/zhenhua32/xingkong/web/app"
)

func main() {
	app.RunSwaggerCmd()

	command := "go run ./cmd/server/main.go"

	fmt.Println("准备启动 Web 服务器, 执行的命令是", command)
	params := strings.Split(command, " ")
	cmd := exec.Command(params[0], params[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	cmd.Start()
	cmd.Wait()
}
