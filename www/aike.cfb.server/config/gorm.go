package config

/**
	gorm配置
 */
type GormConfig struct {

}

/**
	初始化gorm相关配置
 */
func newGormConfig() map[string]GormConfig {
	return  make(map[string]GormConfig)
}
