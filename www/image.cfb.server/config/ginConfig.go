package config

/**
	基本配置
 */
type ginConfig struct {
	// 监听端口
	Port				int
	// 是否调试模式
	IsDebug 		bool
}

/**
	初始化gin相关配置
 */
func newGinConfig() *ginConfig {
	return &ginConfig{
		// 作为默认配置
		Port	:	8081,
	}
}
