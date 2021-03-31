package image

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"image-cfb-server/kernel/base"
	"image-cfb-server/module/image/constant"
)

var UEditorController = new(uEditorController)

type uEditorController struct {
	// 继承
	_ base.BaseController
}

/**
	通过 action 进行业务调用
 */
func (c *uEditorController)Index(context *gin.Context)  {
	// 通过 action 判断业务
	action := context.Query("action")

	fmt.Println("当前 action：", action)

	switch action {
		case constant.UEditorAction_Config:

		case constant.UEditorAction_UploadImage:

	default:

	}

	// 当 action == config 时，返回配置信息

	// 当 action 为 uploadimage 等类型时，进行图片上传

	// 当 action 为 uploadscrawl 等类型时，进行涂鸦上传

	// 当 action 为 uploadvideo、uploadfile 时，进行上传视频以及文件

	// 否则，抛出错误
}
