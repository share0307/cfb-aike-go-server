package aike

import (
	"aike-cfb-server/config"
	"aike-cfb-server/http/router/aike/api"
	"aike-cfb-server/kernel/helper"
	"fmt"
	"github.com/gin-gonic/gin"
)

/**
	启动服务
 */
func Run()  {
	// 设置运行模式
	gin.SetMode(config.Conf.Gin.Mode)

	// 实例化gin
	router := gin.New()
	// 加载html模板
	router.LoadHTMLGlob(fmt.Sprintf("%s/%s/*/*", helper.GetResourcePath(), "tmpl"))

	// api 分组路由
	apiGroup := router.Group("/api")
	// 设置图片分组路由
	api.SetApiImageRouter(apiGroup)

	// 开启 goroutine ，运行服务
	router.Run(config.Conf.Gin.Addr)
}