package config

import "github.com/gin-gonic/gin"

/**
	基本配置
 */
type ginConfig struct {
	// 监听地址
	Addr				string
	// 是否调试模式
	IsDebug 			bool
	// 运行模式
	Mode				string
}

/**
	初始化gin相关配置
 */
func newGinConfig() ginConfig {
	return ginConfig{
		// 作为默认配置
		Addr				:		"0.0.0.0:8081",
		IsDebug				:		false,
		Mode				:		gin.ReleaseMode,
	}
}
