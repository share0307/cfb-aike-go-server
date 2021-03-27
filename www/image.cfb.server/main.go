package main

import (
	"image-cfb-server/cmd"
	"image-cfb-server/provider"
)

func main() {
	// 初始化配置
	initConfig()

	// 初始化服务提供者
	initProvider()

	// 执行命令
	cmd.Execute()
}

/**
	初始化配置
 */
func initConfig()  {
	// 初始化时区

	// 初始化
}

/**
	初始化配置
 */
func initProvider()  {
	// 初始化日志服务提供者
	provider.InitLogger()
}
