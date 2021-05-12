package dao

import (
	"aike-cfb-server/provider"
)

// 此dao中全局使用的db链接
var db *provider.GormProvider = provider.NewGormProvider("default")

/**
	初始化
 */
func init() {

}
