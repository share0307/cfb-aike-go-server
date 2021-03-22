package main

import (
	"github.com/gin-gonic/gin"
	"image-cfb-server/http/router/api"
)

/**
	执行方法
 */
func main()  {
	// 初始化
	router := gin.Default()
	// api 分组
	apiGroup := router.Group("/api", func(context *gin.Context) {
		// 所有此分组的路由，都将走此方法
	})
	{
		// 文件分组
		apiFileGroup := apiGroup.Group("/files");
		{
			file := &api.FIle{}
			// 上传
			apiFileGroup.GET("/show/:id", file.Show)
		}
	}

	router.Run(":8081")
}
