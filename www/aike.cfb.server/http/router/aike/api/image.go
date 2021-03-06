package api

import (
	"aike-cfb-server/http/controller/api/image"
	"github.com/gin-gonic/gin"
)

/**
	设置api路由分组的陆游
 */
func SetApiImageRouter(group *gin.RouterGroup){
	// 图片上传页面
	group.GET("/images/index",  image.CommonController.Index)
}
