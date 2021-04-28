package provider

import "aike-cfb-server/config"

/**
	gorm的服务提供者
 */
type gormProvider struct {
	config config.GormConfig
}

/**
	获取gorm的链接
 */
func (g *gormProvider)Connect()  {

}
