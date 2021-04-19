package config

/**
	基本配置
 */
type baseConfig struct {
	// 是否调试模式
	IsDebug 		bool
}

/**
	初始化配置
 */
func newBaseConfig() baseConfig  {
	return baseConfig{

	}
}