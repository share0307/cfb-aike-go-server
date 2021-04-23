package business;

import (
	"aike-cfb-server/kernel/base"
	"fmt"
)

// 作为单例被调用
var WechatMessageBusiness *wechatMessageBusiness = new(wechatMessageBusiness)

/**
	文件业务
 */
type wechatMessageBusiness struct {
	// 继承
	_ base.BaseBusiness
}

/**
	处理微信消息
 */
func (w *wechatMessageBusiness)HandleWechatMessage(content string)  {
	fmt.Println("接收到消息：", content)
}
