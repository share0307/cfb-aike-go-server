package main

import (
	"image-cfb-server/cmd"
	"image-cfb-server/provider"
)

/**
	应用入口
 */
func main() {
	// 初始化应用
	initApp()

	// 初始化服务提供者
	initProvider()

	// 执行命令
	cmd.Execute()
}

/**
	初始化应用
*/
func initApp() {
	// 初始化时区

	// 初始化
}

/**
初始化配置
*/
func initProvider() {
	// 初始化日志服务提供者
	provider.InitLogger()
}
