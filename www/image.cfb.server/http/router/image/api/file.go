package api

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/http/controller/api/file"
)

/**
设置api路由分组的陆游
*/
func SetApiFileRouter(group *gin.RouterGroup)  {
	// 图片上传页面
	group.GET("/files/index",  file.CommonController.Index)
}
