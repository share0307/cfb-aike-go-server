package message

import (
	"aike-cfb-server/kernel/component/queue"
	"aike-cfb-server/provider"
	"context"
	"fmt"
	"github.com/streadway/amqp"
	"sync"
	"time"
)

/**
	处理微信消息
 */
type HandleWechatMessage struct {
	// 嵌入通用结构体
	queue.CommonQueueImplementation
}

/**s
	启动服务
 */
func (h *HandleWechatMessage)Init(ctx context.Context, group *sync.WaitGroup) {
	provider.LoggerProvider.Info("正在启用 HandleWechatMessage 的队列服务！")

	// 启动监控
	h.Monitor(ctx, group)

	// 设置关闭后调用的函数
	h.SetDownFunc(func() {
		h.Down()
	})

	// 设置别名
	h.SetQueueConfig("default")

	// 初始化队列
	h.InitMq()

	go h.HandlePublishMsgProcess()
}

/**
	关闭服务
 */

func (h *HandleWechatMessage)Down() {
	fmt.Println("HandleWechatMessage 关闭了！！")
}

// 处理消息的流程，从队列中获取消息，会推送到此方法中
func (h *HandleWechatMessage)HandleReceiveMsgProcess(delivery *amqp.Delivery) {
	fmt.Println(delivery.Body)

	delivery.Ack(false)
}

// 发送消息的流程，会从此方法中取得数据，然后推送队列中
func (h *HandleWechatMessage)HandlePublishMsgProcess() {
	// 数据处理

	t := time.NewTicker(1 * time.Second)
	// 数据发送
	for  {
		select {
			case <- t.C:
				fmt.Println("aaaaaa")
				h.PublishSimpleMsg([]byte("hello world：" + time.Now().String()))
		}

	}
	// 成功。。

	// 失败。。
}

// 处理链接异常的流程
func (h *HandleWechatMessage)HandleConnectionErrProcess() {

}

/**
消息出错的流程通用实现
todo：默认不作任何业务
*/
func (h *HandleWechatMessage)HandleMsgErrProcess() {
	return
}