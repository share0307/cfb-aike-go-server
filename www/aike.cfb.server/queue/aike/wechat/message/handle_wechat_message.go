package message

import (
	"aike-cfb-server/kernel/component/queue"
	"aike-cfb-server/provider"
	"context"
	"fmt"
	"sync"
	"time"
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
func (h *HandleWechatMessage)On(ctx context.Context, group *sync.WaitGroup)  {
	provider.LoggerProvider.Info("正在启用 HandleWechatMessage 的队列服务！")

	// 初始化服务
	provider.LoggerProvider.Debugf("准备初始化服务：", time.Now().String())
	h.Init()
	provider.LoggerProvider.Debugf("完成初始化服务：", time.Now().String())

	// 进行监听 context
	select {
		case <- ctx.Done():
			h.Down()
			// 完成分组任务
			group.Done()
	}
}

/**
	关闭服务
 */

func (h *HandleWechatMessage)Down() {
	fmt.Println("HandleWechatMessage 关闭了！！")
}

// 队列相关消息
// 做一些初始化工作
func (h *HandleWechatMessage)Init() {
	// 初始化一些配置

	// 初始化一些通用组件

	// 初始化rabbitmq链接
}

// 处理消息的流程，从队列中获取消息，会推送到此方法中
func (h *HandleWechatMessage)HandleReceiveMsgProcess() {

}

// 发送消息的流程，会从此方法中取得数据，然后推送队列中
func (h *HandleWechatMessage)HandlePublishMsgProcess() {

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