package base

import (
	"github.com/streadway/amqp"
	"sync"
)

// 思路
// 1. 定义好 rabbitmq 的 链接
// 2. 定义好 rabbitmq 的 渠道
// 3. 定义好 生产者
// 4. 定义好消费者

// 需实现的方法：
// 1. 链接/重连机制
// 2. 心跳，查看rabbitmq内部是否有实现
// 3. 获取消息机制
// 4. 发送消息机制
// 5. 重试机制
// 6. 错误队列机制

// 行为 && hook
// 1. 发送消息
// 2. 接收消息
// 3. 错误捕获

/**
	基本 rabbitmq 队列
 */
type RabbitMQ struct {
	// amqp 链接
	connection *amqp.Connection
	// 渠道
	channel *amqp.Channel
	// 交换机名称
	exchangeName string
	// 交换机类型
	exchangeType string
	// 队列名称
	queueName string
	// 路由名称
	routeName string
	//

	// 读写锁
	sync.RWMutex
}

/**
	基本队列，可提供实现业务，考虑做成驱动式
 */
type BaseQueue struct {

}