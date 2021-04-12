package message

import (
	"aike-cfb-server/kernel/component/queue"
	"context"
	"fmt"
	"github.com/go-redis/redis"
)

/**
	处理微信消息
 */
type HandleWechatMessage struct {
	queue.CommonQueueImplementation
}

/**
	启动服务
 */
func (h *HandleWechatMessage)On(ctx context.Context)  {

	fmt.Println("启动 HandleWechatMessage 服务")

	go func() {
		select {
			case <- ctx.Done():
				fmt.Println("关闭了！！", ctx.Err())
				h.Down()
		}
	}()

	fmt.Println("关闭 HandleWechatMessage 服务")
}

/**
	关闭服务
 */
func (h *HandleWechatMessage)Down()  {
	fmt.Println("HandleWechatMessage 关闭了！！")
}

// 队列相关消息
// 做一些初始化工作
func (h *HandleWechatMessage)Init() {

}

/**
获取队列名称，用于生成消息版本号，防止污染的问题
 */
func (h *HandleWechatMessage)GetQueueName() string {

	return ""
}

// 处理消息的流程，从队列中获取消息，会推送到此方法中
func (h *HandleWechatMessage)HandleReceiveMsgProcess() {

}

// 发送消息的流程，会从此方法中取得数据，然后推送队列中
func (h *HandleWechatMessage)HandlePublishMsgProcess() {

}

// 出现异常时的流程
func (h *HandleWechatMessage)HandleMsgErrProcess() {

}

// 处理链接异常的流程
func (h *HandleWechatMessage)HandleConnectionErrProcess() {

}

/**
	生成 sign 的规则
 */
func (h *HandleWechatMessage)GetDuplicateMap() map[string]string {
	return map[string]string{

	}
}

func (h *HandleWechatMessage) SetDuplicateRds(rds *redis.Client) {
	panic("implement me")
}
