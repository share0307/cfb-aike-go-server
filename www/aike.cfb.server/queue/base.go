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
	运行队列
 */
type ExecMqService interface {
	Run(ctx context.Context, group *sync.WaitGroup)
}

/**
	导出业务
 */
func Exec(mqService ExecMqService)  {
	// 任务分组
	waitGroup := &sync.WaitGroup{}

	// 上下文
	ctx,cancel := context.WithCancel(context.Background())

	// 运行mq
	mqService.Run(ctx, waitGroup)

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

	// 的呢古代分组任务完结
	waitGroup.Wait()

	fmt.Println("所有协程已退出，关闭进程！！")
}
