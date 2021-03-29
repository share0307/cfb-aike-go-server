package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"image-cfb-server/kernel/helper"
)

/**
	系统配置结构体
 */
type Config struct {
	// ***************** 可以进行 toml 配置的 *************/
	// 基础配置
	Base		*baseConfig	`mapstructure:"base"`
	// gin框架相关配置
	Gin			*ginConfig	`mapstructure:"gin"`


	// ***************** 不需要进行 toml 配置的 *************/
	Logger		*loggerConfig
}

var (
	// 配置全局单例
	Conf *Config
	// 获取配置目录
	configPath string = helper.GetRelativePathWithPanic("config")
	// vp
	vp *viper.Viper
)

/**
	初始化配置
 */
func init() {
	// 初始化 vp 对象
	initViper()

	// 初始化配置
	initConfig()
}

/**
	初始化 vp 对象
 */
func initViper()  {
	var err error;

	// 若是使用监听配置变更，则必须要手工初始化
	vp = viper.New()

	// 加入配置搜索路径
	vp.AddConfigPath(configPath)
	// 设置使用的配置文件名称
	vp.SetConfigName("config")
	// 设置使用文件配置类型
	vp.SetConfigType("toml")

	// 读取/加载配置
	if err = vp.ReadInConfig();err != nil {
		panic("读取文件配置失败：" + err.Error())
	}
}

/**
	初始化配置，基于viper，toml的配置项，大于初始化时的默认配置
 */
func initConfig()  {
	Conf = &Config{
		Base		:		newBaseConfig(),
		Gin			:		newGinConfig(),
		Logger		:		newLoggerConfig(),
	}

	mapConfigFromToml(Conf)
}

/**
	从toml中加载配置
 */
func mapConfigFromToml(config *Config)  {
	var err error

	// 映射配置
	if err = vp.Unmarshal(config); err != nil {
		panic("viper映射配置失败：" + err.Error())
	}

	// 开启监听配置
	vp.WatchConfig()
	// 配置变更后的毁掉方法
	vp.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件发生变更，将重新加载！")

		if err = vp.Unmarshal(config); err != nil {
			panic("viper映射配置失败：" + err.Error())
		}

		fmt.Println("配置文件，已重新加载完毕！")
	})
}