package helper

import (
	"os"
	"path/filepath"
)

// 获取根目录
var RootPath string = getRootPath()
// 获取
var envName string

/**
	获取根路径
	1. 先从环境变量 CFB_ROOT_PATH 中获取项目根路径，多用于项目生产部署
	2. 当没设置 CFB_ROOT_PATH 时，则取当前路径作为
 */
func getRootPath()  string {
	var rootPath string
	var exists bool
	var  err error

	// 默认先取环境变量
	if rootPath, exists = os.LookupEnv("CFB_ROOT_PATH"); !exists {
		// 不存在则取当前路径
		rootPath,err = filepath.Abs(filepath.Dir(os.Args[0]))

		if err != nil {
			panic("获取项目根目录失败！")
		}
	}

	if _,err := os.Stat(rootPath); err != nil{
		if os.IsNotExist(err) {
			panic("项目根目录不存在："+ err.Error())
		}
	}

	return rootPath
}