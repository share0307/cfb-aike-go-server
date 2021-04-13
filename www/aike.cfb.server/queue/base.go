package queue

import (
	"aike-cfb-server/kernel/component/queue"
	"context"
	"fmt"
	"os"
	"os/signal"
)

/**
	队列接口
 */
type ExportQueueInterface interface {
	// 此接口，也必须要实现这些
	queue.QueueInterface

	// 启动服务
	On(ctx context.Context)

	// 关闭服务时，做一些清理的工作
	// 说明：此方法更多是侧重作一些资源清理的事情，而On中的context更多是为了关联goroutine上下文
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

	// 常驻进程
	//select {}

	// 监听信号
	c := make(chan os.Signal)
	// 监听所有信号！
	signal.Notify(c)
	//监听指定信号
	//signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2,syscall.SIGINT)

	//阻塞直至有信号传入
	s := <-c
	fmt.Println("捕获到信号：", s)
}
