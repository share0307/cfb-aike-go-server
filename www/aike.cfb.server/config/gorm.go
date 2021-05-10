package config

/**
	gorm配置
 */
type GormConfig struct {
	// 链接信息
	Dsn string
}

/**
	初始化gorm相关配置
 */
func newGormConfig() map[string]GormConfig {
	return  make(map[string]GormConfig)
}
