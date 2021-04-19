package queue

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

/**
	队列接口
 */
type ExportQueueInterface interface {

	// 启动服务
	On(ctx context.Context, group *sync.WaitGroup)

	// 关闭服务时，做一些清理的工作
	// 说明：此方法更多是侧重作一些资源清理的事情，而On中的context更多是为了关联goroutine上下文
	Down()
}

/**
	导出业务
 */
func Consume(queues []ExportQueueInterface)  {
	// 生成context
	ctx, cancel := context.WithCancel(context.Background())
	// 取消，必须注意：肯定会调用此 defer，但是协程是否来得及处理，这个必须得靠手工保证！！
	//defer cancel()

	// 任务分组
	waitGroup := &sync.WaitGroup{}

	// 启动goroutine
	for _,queue := range queues {
		waitGroup.Add(1)
		// 启动协程
		go queue.On(ctx, waitGroup)
	}

	// 常驻进程
	//select {}

	// 监听信号
	c := make(chan os.Signal)
	// 监听所有信号！
	//signal.Notify(c)
	//监听指定信号
	signal.Notify(c, syscall.SIGHUP, syscall.SIGUSR2,syscall.SIGINT)

	//阻塞直至有信号传入
	s := <-c
	fmt.Println("捕获到信号：", s)
	fmt.Println("等待协程退出！！")

	// 通知协程退出！！
	cancel()

	waitGroup.Wait()

	fmt.Println("所有协程已退出，关闭进程！！")
}
