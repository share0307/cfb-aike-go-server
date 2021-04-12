package wechat

import (
	"aike-cfb-server/queue"
	"aike-cfb-server/queue/aike/wechat/message"
)

func On()  {
	// 定义服务队列
	queues  := []queue.ExportQueueInterface{
		// 处理微信消息
		&message.HandleWechatMessage{},
	}


	// 启动goroutine
	queue.Export(queues)
}
