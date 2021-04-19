package config

/**
	rabbitmq的配置
 */
type RabbitmqConfig struct {
	// 链接信息
	Dsn string
	// qos 服务质量
	Qos int
	// 心跳
	Heartbeat int
	// 渠道
	Channel string
	// 队列
	Queue string
}

/**
	rabbitmq相关配置
 */
func newRabbitmqConfig() map[string]RabbitmqConfig {
	return  make(map[string]RabbitmqConfig)
}
