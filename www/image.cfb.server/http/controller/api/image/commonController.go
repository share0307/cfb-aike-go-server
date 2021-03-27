package image

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/kernel/base"
)

var CommonController = new(commonController)

type commonController struct {
	// 继承
	_ base.BaseController
}

/**
	首页
 */
func (c *commonController)Index(context *gin.Context)  {
	context.Writer.WriteString("hello world!")
}
