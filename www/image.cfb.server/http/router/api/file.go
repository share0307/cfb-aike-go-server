package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

type FIle struct {

}

/**
	文件上传
 */
func (file *FIle)Upload(context *gin.Context)  {
	context.Writer.WriteString("upload!!")
}

/**
	文件显示
 */
func (file *FIle)Show(context *gin.Context)  {
	fileId,_ := context.Params.Get("id")

	fmt.Println("file_id:", fileId)
}
