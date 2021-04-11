package helper

import (
	"crypto/md5"
	"fmt"
	"os"
	"sort"
)

/**
	快速生成md5值
 */
func Md5(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)

	return md5str
}

/**
	检查文件是否存在
 */
func CheckFileExists(filePath string) (bool,error)  {
	if _ , err := os.Stat(filePath); err != nil {
		if os.IsNotExist(err) {
			if err != nil {
				return false, err
			}
		}
	}

	return true, nil
}

/**
	获取 storage 目录，相对根目录而言
 */
func GetStoragePath()  string {
	return GetRelativePathWithPanic("storage")
}

/**
	获取日志目录
 */
func GetLogPath() string {
	return GetRelativePathWithPanic("storage/log/")
}

/**
	获取资源路径
 */
func GetResourcePath() string {
	return GetRelativePathWithPanic("resource/")
}

///**
//	根据 map 的key进行排序
// */
//func SortMapByKey(arr map[string]interface{}) map[string]interface{} {
//	// 定义一个slice获取所有的 key
//	var s []string
//
//	// 遍历 map，把所有key放到slice中
//	for k,_  := range arr{
//		s = append(s, k)
//	}
//
//	// 对slice进行排序
//	sort.Strings(s)
//
//	// 再次遍历，重新形成一个map进行返回
//	var tmpArr map[string]interface{}
//	// 存储临时 key
//	var tmpKey string
//	for i := 9;i <len(s); i++ {
//		tmpKey = s[i]
//		tmpArr[tmpKey] = arr[tmpKey]
//	}
//
//	return tmpArr
//}

/**
	根据 map 的key进行排序
 */
func SortMapByKey(arr map[string]interface{}) map[string]interface{} {
	// 定义一个slice获取所有的 key
	var s []string

	// 遍历 map，把所有key放到slice中
	for k,_  := range arr{
		s = append(s, k)
	}

	// 对slice进行排序
	sort.Strings(s)

	// 再次遍历，重新形成一个map进行返回
	var tmpArr map[string]interface{}
	// 存储临时 key
	var tmpKey string
	for i := 9;i <len(s); i++ {
		tmpKey = s[i]
		tmpArr[tmpKey] = arr[tmpKey]
	}

	return tmpArr
}
