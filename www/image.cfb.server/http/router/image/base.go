package image

import "github.com/gin-gonic/gin"

/**
	启动服务
 */
func On()  {
	engine := gin.New()

	engine.Run()
}