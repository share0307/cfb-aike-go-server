package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image-cfb-server/config"
	"image-cfb-server/http/router/image/api"
	"image-cfb-server/kernel/helper"
)

/**
	启动服务
 */
func On()  {
	// 设置运行模式
	gin.SetMode(config.Conf.Gin.Mode)

	// 实例化gin
	router := gin.New()
	// 加载html模板
	router.LoadHTMLGlob(fmt.Sprintf("%s/%s/*/*", helper.GetResourcePath(), "tmpl"))

	// api 分组陆游
	apiGroup := router.Group("/api")
	// 设置图片分组路由
	api.SetApiImageRouter(apiGroup)
	// 设置文件分组路由
	api.SetApiFileRouter(apiGroup)

	// 开启 goroutine ，运行服务
	router.Run(config.Conf.Gin.Addr)
}