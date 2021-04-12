package helper

import (
	"strconv"
	"testing"
)

/**
	测试md5的方法
 */
func TestMd5(t *testing.T) {
		if Md5("12345678") == Md5("87654321") {
			t.Error("错误！！")
		}
}

/**
	测试md5的方法
 */
func BenchmarkMd5(b *testing.B) {
		for i  := 0;i <b.N; i++ {
			Md5(strconv.Itoa(i))
			//fmt.Println(tmp)
		}
}