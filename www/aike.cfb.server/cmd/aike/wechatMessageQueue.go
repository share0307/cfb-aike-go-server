package aike

import (
	"aike-cfb-server/queue/aike/wechat"
	"github.com/spf13/cobra"
)

var WechatMessageQueue = &cobra.Command{
	Use:   "aike/wechatMessageQueue",
	Short: "消费艾克的微信数据",
	Long: `处理微信消息队列`,
	Run: func(cmd *cobra.Command, args []string) {
		// 启动服务，启动 web server 服务
		wechat.Run()
	},
}