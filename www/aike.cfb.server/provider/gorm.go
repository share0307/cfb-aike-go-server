package provider

import (
	"aike-cfb-server/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"time"
)

/**
	实例化一个Gorm的服务提供者
 */
func NewGormProvider(alias string) *gormProvider {
	// 获取配置
	gormConfig,exists := config.Conf.Gorms[alias]

	if !exists {
		panic("mq配置不存在！")
	}

	gProvider := &gormProvider{
		config: gormConfig,
	}

	return gProvider
}

/**
	gorm的服务提供者
	当前只支持mysql
 */
type gormProvider struct {
	config 		config.GormConfig
	DB				*gorm.DB
}

/**
	获取gorm的链接
 */
func (g *gormProvider)InitGorm()  {
	// 准备参数

	// 链接
	g.Connect()

	// 开启监听
	go g.monitor()
}

func (g *gormProvider)Connect()  {
	db, err := gorm.Open("mysql", g.config.Dsn)

	if err != nil {
		panic("链接数据库失败！原因为：" + err.Error())
	}

	// 设置连接池等

	g.DB = db
}

func (g *gormProvider)monitor()  {
	t := time.NewTicker(3 * time.Second)

	for  {
		select {
			case <- t.C:
				if err := g.DB.DB().Ping(); err != nil {
					LoggerProvider.Error("数据库无法ping通，原因为：" + err.Error())

					// 进行重新链接
					g.Connect()
				}
		}
	}
}
