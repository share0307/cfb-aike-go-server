package aike

import (
	"aike-cfb-server/queue/aike"
	"github.com/spf13/cobra"
)

var WechatQueueCmd = &cobra.Command{
	Use:   "aike/wechatQueue",
	Short: "消费艾克的微信数据",
	Long: `处理微信消息队列`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动服务，启动 web server 服务
		aike.Run()
	},
}