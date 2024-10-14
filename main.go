package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	// Docker 和 docker-compose 的完整路径
	dockerPath := ""
	dockerComposePath := ""

	// 获取命令行参数
	args := os.Args

	// 如果输入的命令是 `docker compose`，则重定向到 `docker-compose`
	if len(args) > 1 && args[1] == "compose" {
		// 将 "docker compose" 替换为 "docker-compose" 并传递后续参数
		composeArgs := args[2:] // 直接传递 "compose" 后面的参数
		cmd := exec.Command(dockerComposePath, composeArgs...)
		
		// 将命令的输出和错误信息传递到当前进程
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// 执行命令
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing docker-compose: %v\n", err)
			os.Exit(1)
		}
	} else {
		// 如果不是 `docker compose`，则直接调用 `docker` 命令
		cmd := exec.Command(dockerPath, args[1:]...)
		
		// 将命令的输出和错误信息传递到当前进程
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
		cmd.Stdin = os.Stdin

		// 执行命令
		err := cmd.Run()
		if err != nil {
			fmt.Printf("Error executing docker: %v\n", err)
			os.Exit(1)
		}
	}
}
