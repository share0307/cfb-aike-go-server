package aike

import (
	"aike-cfb-server/kernel/component/queue"
	mq "aike-cfb-server/queue"
	"aike-cfb-server/queue/aike/wechat/message"
)

/**
	运行服务
 */
func Run()  {
	// 实例化mq服务容器
	mqService := queue.NewMqService()

	// 添加接受的微信消息队列
	mqService.Add(new(message.HandleWechatMessage))

	// 通用运行
	mq.Exec(mqService)
}

