package helper

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
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

	if exists,err := CheckFileExists(rootPath); !exists {
		panic("项目根目录不存在："+ err.Error())
	}

	return rootPath
}

/**
	获取相对根目录的目录
 */
func GetRelativePath(relativePath string)  (string, error) {
	// 去掉两边空格
	relativePath = strings.TrimSpace(relativePath)
	// 统一分隔符
	relativePath = strings.ReplaceAll(relativePath, "\\", "/")
	// 去掉左右两边的 /
	relativePath = strings.Trim(relativePath, "/")

	// 组装相对路径
	var absPath string = fmt.Sprintf("%s/%s", RootPath, relativePath)

	if  exists,err := CheckFileExists(absPath) ;!exists {
		return "", err
	}

	return absPath, nil
}

/**
	获取相对根目录的目录，当文件不存在时，直接触发 panic
 */
func GetRelativePathWithPanic(relativePath string)  string {
	absPath, err := GetRelativePath(relativePath)

	if err != nil {
		panic(err)
	}

	return absPath
}