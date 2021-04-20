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
func NewRabbitmqProvider(alias string)  *RabbitmqProvider {
	// 获取配置
	mqConfig,exists := config.Conf.RabbitMqs[alias]

	if !exists {
		panic("mq配置不存在！")
	}

	// 返回mq实例
	mqProvider := &RabbitmqProvider{
		Config: mqConfig,
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
	Config config.RabbitmqConfig
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
		Heartbeat	:		time.Duration(r.Config.Heartbeat) * time.Second,
		Vhost				: 		r.Config.Vhost,
	}

	r.connect,err = amqp.DialConfig(r.Config.Dsn, amqpConfig)

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
		r.Config.Exchange,
		r.Config.ExchangeType,
		true,
		false,
		false,
		false,
		nil,
		)

	// 声明队列
	r.channel.QueueDeclare(
		r.Config.Queue,
		true,
		false,
		false,
		false,
		nil,
		)

	// 帮顶队列与交换机的关系
	r.channel.QueueBind(
		r.Config.Queue,
		r.Config.Route,
		r.Config.Exchange,
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

	err := r.channel.Publish(r.Config.Exchange, r.Config.Route, false, false, msg)

	return err
}

/**
	检测是否lianjiezhong
 */
func (r *RabbitmqProvider)IsClose() bool {
	return r.connect.IsClosed()
}
