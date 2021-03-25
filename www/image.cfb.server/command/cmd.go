package command

import (
	"fmt"
	"github.com/spf13/cobra"
)

// 记录全部命令
var RootCmd *cobra.Command = new(cobra.Command)

/**
	初始化所有命令
 */
func Initialize() {
	// 把所有命令均添加为子命令
	for _,cmd := range ExportCommands(){
		RootCmd.AddCommand(cmd)
	}
}

/**
	运行
 */
func Execute() {
	RootCmd.Execute()
}

/**
	导出子命令
 */
func ExportCommands() []*cobra.Command {
	var cmds []*cobra.Command = []*cobra.Command {
		&cobra.Command{
			Use		:		"test",
			Run	: 		func(cmd *cobra.Command, args []string) {
				fmt.Println("hello world!")
			},
		},
	}

	return cmds;
}

