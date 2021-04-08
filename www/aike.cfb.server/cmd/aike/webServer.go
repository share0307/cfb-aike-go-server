package aike

import (
	"aike-cfb-server/http/router/aike"
	"github.com/spf13/cobra"
)

var WebServerCmd = &cobra.Command{
	Use:   "aike/webServer",
	Short: "艾客 http 服务",
	Long: `提供艾客 web http server，提供下载/显示等功能`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动服务，启动 web server 服务
		aike.On()
	},
}