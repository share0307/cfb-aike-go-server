package wechat

import (
	"aike-cfb-server/queue"
	"aike-cfb-server/queue/aike/wechat/message"
)

func Run()  {
	// 定义服务队列
	queues  := []queue.ExportQueueInterface{
		// 处理微信消息
		&message.HandleWechatMessage{},
	}

	// 启动 goroutine
	//queue.Export(queues)

	// 设置队列

	// 进行绑定
	// todo：1. 为什么一定要设置好 "粘合剂" 呢？因为在golang中，父类是无法直接调用 子类的方法(哪怕是重写)
	// todo：2. 所以一些通用组件都在父类中实现的话，一些子类的业务，就必须要通过 "粘合剂" 这个角色去进行把数据流绑定到一起了
}

