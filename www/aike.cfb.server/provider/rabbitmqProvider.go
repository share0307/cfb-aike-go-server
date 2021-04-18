package provider

import (
	"aike-cfb-server/config"
	"github.com/streadway/amqp"
	"time"
)

/**
	实例化rabbitmq的服务提供者
 */
func NewRabbitmqProvider(alias string)  *rabbitmqProvider {
	// 获取配置
	mqConfig,exists := config.Conf.RabbitMqs[alias]

	if !exists {
		panic("mq配置不存在！")
	}

	// 返回mq实例
	return &rabbitmqProvider{
		Config: mqConfig,
	}
}

/**
	rabbitmq的服务提供者
 */
type rabbitmqProvider struct {
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
func (r *rabbitmqProvider)Connect()  {
	var err error
	var amqpConfig = amqp.Config{
		Heartbeat: time.Duration(r.Config.Heartbeat) * time.Second,
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
	r.channel.Qos()
}
