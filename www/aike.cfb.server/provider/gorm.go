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
func NewGormProvider(alias string) *GormProvider {
	// 获取配置
	gormConfig,exists := config.Conf.Gorms[alias]

	if !exists {
		panic("gorm配置不存在！")
	}

	gProvider := &GormProvider{
		config: gormConfig,
	}

	// 初始化gorm
	gProvider.InitGorm()

	return gProvider
}

/**
	gorm的服务提供者
	当前只支持mysql
 */
type GormProvider struct {
	// 配置
	config 		config.GormConfig
	// 提供服务
	Server   	*gorm.DB
}

/**
	获取gorm的链接
 */
func (g *GormProvider)InitGorm()  {
	// 准备参数

	// 链接
	g.Connect()

	// 开启监听
	go g.monitorAndReset()
}

/**
	数据库链接
 */
func (g *GormProvider)Connect()  {
	db, err := gorm.Open("mysql", g.config.Dsn)

	if err != nil {
		panic("链接数据库失败！原因为：" + err.Error())
	}

	// 设置连接池等
	// SetMaxIdleConns：设置空闲链接吃鱼中的最大链接数
	db.DB().SetMaxIdleConns(10)

	// SetConnMaxIdleTime：链接齿里面的链接最大空闲时长
	// 当链接持续空闲市场达到 maxIdleTime 后，该链接就会被关闭并重链接池移除，哪怕当前空闲链接数已经小于 SetMaxIdleConns(maxIdleConns)设置的值
	// 链接每次被使用后，持续空闲时长会被重置，从0开始重新计算
	db.DB().SetConnMaxIdleTime(3600 * time.Second)

	// SetMaxOpenConns：设置数据库链接最大打开数
	db.DB().SetMaxOpenConns(100)

	// SetConnMaxLifeTime：设置可重用链接的最长时间
	db.DB().SetConnMaxLifetime(60 * time.Second)

	g.Server = db
}

/**
	对 链接 继续宁监控，若出现链接断开，则进行重连
 */
func (g *GormProvider)monitorAndReset()  {
	t := time.NewTicker(3 * time.Second)

	for  {
		select {
			case <- t.C:
				if err := g.Server.DB().Ping(); err != nil {
					LoggerProvider.Error("数据库无法ping通，原因为：" + err.Error())

					// 进行重新链接
					g.Connect()
				}
		}
	}
}
