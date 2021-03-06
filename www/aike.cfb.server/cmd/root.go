/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this aike except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"aike-cfb-server/cmd/aike"
	"fmt"
	"github.com/spf13/cobra"
	"os"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var cfgFile string

// RootCmd represents the base command when called without any subcommands
// 初始化根命令
var RootCmd = &cobra.Command{
//	Use:   "main",
//	Short: "A brief description of your application",
//	Long: `A longer description that spans multiple lines and likely contains
//examples and usage of using your application. For example:
//
//Cobra is a CLI library for Go that empowers applications.
//This application is a tool to generate the needed files
//to quickly create a Cobra application.`,
//	// Uncomment the following line if your bare application
//	// has an action associated with it:
//	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the RootCmd.
// 执行所有命令
func Execute() {
	cobra.CheckErr(RootCmd.Execute())
}

// 初始化命令
func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config aike (default is $HOME/.main.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	RootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")

	// 初始化命令
	initCommand()
}

// initConfig reads in config aike and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config aike from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		cobra.CheckErr(err)

		// Search config in home directory with name ".main" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".main")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config aike is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Fprintln(os.Stderr, "Using config aike:", viper.ConfigFileUsed())
	}
}

/**
	初始化子命令
 */
func initCommand()  {
	// 艾客 web 服务
	RootCmd.AddCommand(aike.WebServerCmd)
	// 处理微信消息的队列
	RootCmd.AddCommand(aike.WechatQueueCmd)
	// 艾克数据迁移
	RootCmd.AddCommand(aike.MigrateCmd)
}
