package image

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/config"
	"image-cfb-server/http/router/image/api"
)

/**
	启动服务
 */
func On()  {
	// 设置运行模式
	gin.SetMode(config.Conf.Gin.Mode)

	// 实例化gin
	engine := gin.New()

	// api 分组陆游
	apiGroup := engine.Group("/api")
	// 设置图片分组陆游
	api.SetApiImageRouter(apiGroup)
	// 设置文件分组陆游
	api.SetApiFileRouter(apiGroup)

	engine.Run(config.Conf.Gin.Addr)
}