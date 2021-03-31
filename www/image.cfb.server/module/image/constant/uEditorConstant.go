package constant

import "image-cfb-server/kernel/base"

/**
	提供外部使用
 */
var UEditorActionConstant = new(uEditorActionConstant)

/**
	定义枚举值
 */
const (
	// 获取配置
	UEditorAction_Config		= "config"
	// 上传图片
	UEditorAction_UploadImage	= "uploadimage"
	// 涂鸦上传
	UEditorAction_UploadScrawl	= "uploadscrawl"
	// 上传视频
	UEditorAction_UploadVideo	= "uploadvideo"
	// 上传文件
	UEditorAction_UploadFile	= "uploadfile"
)

/**
	UE编辑器动作枚举值
 */
type uEditorActionConstant struct {
	// 继承
	base.BaseConstant
}

/**
	映射枚举
 */
func (c *uEditorActionConstant)GetNames() map[string]interface{} {
	return map[string]interface{}{
		UEditorAction_Config			:	"获取配置",
		UEditorAction_UploadImage		:	"上传图片",
		UEditorAction_UploadScrawl		:	"上传涂鸦",
		UEditorAction_UploadVideo		:	"上传视频",
		UEditorAction_UploadFile		:	"上传文件",
	}
}
