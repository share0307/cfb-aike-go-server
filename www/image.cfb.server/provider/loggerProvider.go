package provider

import (
	"github.com/phachon/go-logger"
	"image-cfb-server/config"
)

var (
	// 提供日志服务
	LoggerProvider *go_logger.Logger
)

func init()  {
	
}

/**
	初始化日志
 */
func InitLogger() {

	LoggerProvider = go_logger.NewLogger()

	// 日志写入对象
	loggerTarget := config.Conf.Logger.Target
	// 配置
	var loggerConfig go_logger.Config

	if loggerTarget == "file"{
		loggerTarget = "file"
		loggerConfig = config.Conf.Logger.Setting.File
	}else{
		loggerTarget = "console"
		loggerConfig = config.Conf.Logger.Setting.Console

		// 因为默认是 console，所以必须要先移除一下原有的中断
		_ = LoggerProvider.Detach(loggerTarget)
	}

	err := LoggerProvider.Attach(config.Conf.Logger.Target, go_logger.LOGGER_LEVEL_DEBUG, loggerConfig)

	if err != nil {
		panic("log 服务提供者初始化打开失败:" + err.Error())
	}
}

