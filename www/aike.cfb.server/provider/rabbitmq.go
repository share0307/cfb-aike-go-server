package provider

import (
	"aike-cfb-server/config"
	"fmt"
	"github.com/streadway/amqp"
	"time"
)

/**
	实例化rabbitmq的服务提供者
 */
func NewRabbitmqProvider(alias string) *RabbitmqProvider {
	// 获取配置
	mqConfig,exists := config.Conf.RabbitMqs[alias]

	if !exists {
		panic("mq配置不存在！")
	}

	// 返回mq实例
	mqProvider := &RabbitmqProvider{
		config	: mqConfig,
	}

	// 启动监听
	go func() {
		t := time.NewTicker(1 * time.Second)

		for  {
			select {
				case <- t.C:
					fmt.Println("监听中")
					if mqProvider.IsClose() == true {
						fmt.Println("已断开！")
					}
			}
		}

	}()

	return mqProvider
}

/**
	rabbitmq的服务提供者
 */
type RabbitmqProvider struct {
	// 保存使用的配置别名
	config config.RabbitmqConfig
	// 保存rabbitmq的链接
	connect *amqp.Connection
	// 保存rabbitmq的渠道
	channel *amqp.Channel
}

/**
	链接
 */
func (r *RabbitmqProvider)Connect()  {
	var err error
	var amqpConfig = amqp.Config{
		Heartbeat	:		time.Duration(r.config.Heartbeat) * time.Second,
		Vhost				: 		r.config.Vhost,
	}

	r.connect,err = amqp.DialConfig(r.config.Dsn, amqpConfig)

	if err != nil {
		panic("链接mq失败，原因为：" + err.Error())
	}

	r.channel, err = r.connect.Channel()

	if err != nil {
		panic("打开mq的channel失败，原因为：" + err.Error())
	}

	// 链接设置
	r.channel.Qos(30, 0, false)
}

/**
	初始化交换机与队列
 */
func (r *RabbitmqProvider)InitExchangeAndQueue() {
	// 声明交换机
	r.channel.ExchangeDeclare(
		r.config.Exchange,
		r.config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
		)

	// 声明队列
	r.channel.QueueDeclare(
		r.config.Queue,
		true,
		false,
		false,
		false,
		nil,
		)

	// 帮顶队列与交换机的关系
	r.channel.QueueBind(
		r.config.Queue,
		r.config.Route,
		r.config.Exchange,
		true,
		nil,
		)
}

/**
	发布消息
 */
func (r *RabbitmqProvider)Publish(msg amqp.Publishing) error {
	// mandatory：当 mandatory 参数为 true 时，交换机无法根据自身和路由键找到一个符合条件的队列，那么 rabbitmq 会调用 Basic.Return 命令，将消息返回给生产者，
	//                           当 mandatory 参数为 false 时，出现上述情形，则消息直接被丢弃
	// immediate：当immediate参数设为true时，如果交换器在将消息路由到队列时发现队列上并不存在任何消费者，那么这条消息将不会存入队列中。当与路由键匹配的所有队列都没有消费者时，该消息会通过Basic.Return返回至生产者。

	err := r.channel.Publish(r.config.Exchange, r.config.Route, false, false, msg)

	return err
}

/**
	消费数据
 */
func (r *RabbitmqProvider)Consume() (<- chan amqp.Delivery, error) {
	deliveryChan,err := r.channel.Consume(
		r.config.Queue,
		"",
		false,		// 是否自动应答
		false,		// exclusive：是否排外的，有两个作用，一：当连接关闭时connection.close()该队列是否会自动删除；二：该队列是否是私有的private，如果不是排外的，可以使用两个消费者都访问同一个队列，没有任何问题，如果是排外的，会对当前队列加锁，其他通道channel是不能访问的，如果强制访问会报异常：com.rabbitmq.client.ShutdownSignalException: channel error; protocol method: #method<channel.close>(reply-code=405, reply-text=RESOURCE_LOCKED - cannot obtain
		false,		// 设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false,		// nowait：是否等待服务器返回
		nil,
		)

	if err != nil {
		panic(err)
	}

	return deliveryChan, err
}

/**
	检测是否lianjiezhong
 */
func (r *RabbitmqProvider)IsClose() bool {
	return r.connect.IsClosed()
}
