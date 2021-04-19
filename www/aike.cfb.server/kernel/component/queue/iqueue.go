package queue

import (
	"context"
	"github.com/go-redis/redis"
	"sync"
)

/**
	定义使用队列时，必须要实现的方法
 */
type IQueue interface {
	// 队列相关消息
	// 做一些初始化工作
	Init(ctx context.Context, group *sync.WaitGroup)
	// 设置队列的别名
	SetQueueConfig(alias string)

	// 业务流程处理
	// 处理消息的流程，从队列中获取消息，会推送到此方法中
	HandleReceiveMsgProcess()
	// 发送消息的流程，会从此方法中取得数据，然后推送队列中
	HandlePublishMsgProcess()
	// 出现异常时的流程
	HandleMsgErrProcess()
	// 处理链接异常的流程
	HandleConnectionErrProcess()

	// 消息去重规则
	// 设置是否开启去重
	SetEnableDuplicateCheckFlag(flag bool)
	// 是否启用去重
	IsEnableDuplicateCheck() bool
	// 去重的规则
	GetDuplicateSign() string
	// 获取 sign 的规则
	GetDuplicateMap() map[string]string
	// 获取 sign 的规则
	SetDuplicateMap(duplicateMap map[string]string)
	// X 秒之内不得重复，也就是说重复 key 的生命周期，单位为秒
	SetDuplicateLifeCycle(second int)
	// X 秒之内不得重复，也就是说重复 key 的生命周期
	GetDuplicateLifeCycle() int
	// 设置重复判断的中间件，依赖redis
	SetRds(rds *redis.Client)
	// 获取重复判断的中间件，依赖redis
	GetRds() *redis.Client

	// 定义一些依赖组建
	// 日志组件，用全局的就好，初始化时必须会初始化全局的日志组件
	//SetLogger(logger *go_logger.Logger)
	////获取日志组件，用于内部写日志
	//GetLogger() *go_logger.Logger
}

