package queue

import (
	"aike-cfb-server/kernel/component/queue"
	"context"
)

/**
	队列接口
 */
type ExportQueueInterface interface {
	// 此接口，也必须要实现这些
	queue.QueueInterface

	On(ctx context.Context)
	Down()
}

/**
	导出业务
 */
func Export(queues []ExportQueueInterface)  {
	// 生成context
	ctx, cancel := context.WithCancel(context.Background())
	// 取消
	defer cancel()

	// 启动goroutine
	for _,queue := range queues {
		// 启动协程
		go queue.On(ctx)
	}

	//time.AfterFunc(3 * time.Second, func() {
	//	cancel()
	//})

	// 常驻进程
	select {}
}
