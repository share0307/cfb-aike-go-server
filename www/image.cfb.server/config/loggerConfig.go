package config

import (
	"fmt"
	go_logger "github.com/phachon/go-logger"
	"image-cfb-server/kernel/helper"
)

type loggerConfig struct {
	// 目标
	Target		string
	// 日志级别
	Level		int
	// 设置
	Setting		struct{
		// 重置台配置
		Console *go_logger.ConsoleConfig
		// 文件配置
		File *go_logger.FileConfig
	}
}

/**
映射、初始化日志配置
*/
func newLoggerConfig() *loggerConfig {

	return &loggerConfig{
		Target	:	vp.GetString("log.target"),
		Setting	:	struct {	// 匿名结构体
			// 重置台配置
			Console *go_logger.ConsoleConfig
			// 文件配置
			File *go_logger.FileConfig
		}{
			// 重置台配置
			Console	: 	&go_logger.ConsoleConfig{
				Color: true,
				JsonFormat: true,
				Format: "",
			},
			// 文件配置
			File	:	&go_logger.FileConfig{
				Filename: fmt.Sprintf("%s/%s/", helper.GetLogPath(), "im.log"),
				LevelFileName: map[int]string{
					go_logger.LOGGER_LEVEL_DEBUG	: fmt.Sprintf("%s/%s", helper.GetLogPath(), "im-debug.log"),
					go_logger.LOGGER_LEVEL_NOTICE	: fmt.Sprintf("%s/%s/", helper.GetLogPath(), "im-notice.log"),
					go_logger.LOGGER_LEVEL_INFO		: fmt.Sprintf("%s/%s/", helper.GetLogPath(), "im-info.log"),
					go_logger.LOGGER_LEVEL_WARNING	: fmt.Sprintf("%s/%s/", helper.GetLogPath(), "im-warning.log"),
					go_logger.LOGGER_LEVEL_ERROR	: fmt.Sprintf("%s/%s/", helper.GetLogPath(), "im-error.log"),
				},
				JsonFormat: true,
				DateSlice: "d",
			},
		},
	}
}
