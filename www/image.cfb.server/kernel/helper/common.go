package helper

import (
	"crypto/md5"
	"fmt"
	"os"
)

var H Helper = Helper{}

type Helper struct {

}

/**
	快速生成md5值
 */
func (helper *Helper)Md5(str string) string{
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