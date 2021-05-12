package image

import (
	"aike-cfb-server/kernel/base"
	"aike-cfb-server/kernel/helper"
	"aike-cfb-server/module/aike/dao"
	"aike-cfb-server/provider"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var CommonController = new(commonController)

type commonController struct {
	// 继承
	_ base.BaseController
}

/**
	首页，用于上传
 */
func (c *commonController)Index(context *gin.Context)  {

	ud := dao.UserDao{}

	ud.Sore()

	ud.Find()

	context.HTML(http.StatusOK, "upload.tmpl", nil)
}

/**
	接收文件上传
 */
func (c *commonController)Upload(context *gin.Context)  {
	file,err := context.FormFile("imgfile")

	if err != nil {
		provider.LoggerProvider.Error("获取上传文件失败！！")
	}

	fmt.Println(context.SaveUploadedFile(file, fmt.Sprintf("%s/%s",helper.GetLogPath(), "aaa.jpeg")))
}
