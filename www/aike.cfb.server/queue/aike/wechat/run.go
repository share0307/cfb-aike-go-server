package wechat

import (
	queueBase "aike-cfb-server/kernel/component/queue"
	"aike-cfb-server/queue"
	"aike-cfb-server/queue/aike/wechat/message"
)

func Run()  {

	qb := queueBase.QueueBind{}
	qb.SetQueue(&message.HandleWechatMessage{})

	qb.Bind()

	// 定义服务队列
	queues  := []queue.ExportQueueInterface{
		// 处理微信消息
		&message.HandleWechatMessage{},
	}

	// 启动 goroutine
	//queue.Export(queues)

	// 设置队列

}

