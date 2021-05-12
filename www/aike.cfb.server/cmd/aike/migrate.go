package aike

import (
	"aike-cfb-server/module/aike/model"
	"aike-cfb-server/provider"
	"github.com/spf13/cobra"
)

var MigrateCmd = &cobra.Command{
	Use:   "aike/migrate",
	Short: "艾客 数据歉意",
	Long: `提供艾客 web http server`,
	Run: func(cmd *cobra.Command, args []string) {
		// 直接执行数据迁移
		// 先链接数据库
		db := provider.NewGormProvider("default")

		db.Server.AutoMigrate(model.UserModel{})
	},
}