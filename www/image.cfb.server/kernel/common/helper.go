package common

import (
	"crypto/md5"
	"fmt"
)

var H Helper = Helper{}

type Helper struct {

}

func (helper *Helper)Md5(str string) string{
	data := []byte(str)
	has := md5.Sum(data)
	md5str := fmt.Sprintf("%x", has)

	return md5str
}