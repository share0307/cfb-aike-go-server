package api

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/http/controller/api/image"
)

/**
	设置api路由分组的陆游
 */
func SetApiImageRouter(group *gin.RouterGroup)  {
	// 图片上传页面
	group.GET("/images/index",  image.CommonController.Index)
	// 接收文件上传
	group.POST("/images/upload",  image.CommonController.Upload)




	// UE编辑器，上传图片
	group.GET("/ueditor/", image.UEditorController.Index)
	//UE编辑器，上传图片
	group.GET("/ueditor/index", image.UEditorController.Index)
	// UE编辑器，上传图片
	group.GET("/ueditor/upload", image.UEditorController.Upload)
}
