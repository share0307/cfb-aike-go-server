package image

import (
	"github.com/spf13/cobra"
	"image-cfb-server/http/router/image"
)

var WebServerCmd = &cobra.Command{
	Use:   "image/webServer",
	Short: "图片服务",
	Long: `提供图片web http server，提供下载/显示等功能`,
	Run: func(cmd *cobra.Command, args []string) {

		// 启动服务
		image.On()
	},
}