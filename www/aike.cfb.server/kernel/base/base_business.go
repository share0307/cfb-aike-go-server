package base

import "fmt"

/**
	模型积累
 */
type BaseBusiness struct {

}

/**
	测试方法
 */
func (b *BaseBusiness)Say(str string)  {
	fmt.Println(str)
}
