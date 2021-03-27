package image

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/config"
)

/**
	启动服务
 */
func On()  {
	engine := gin.New()
	gin.SetMode(config.Conf.Gin.Mode)


	engine.Run(config.Conf.Gin.Addr)
}