package business;

import "aike-cfb-server/kernel/base"

/**
	文件业务
 */
type imageBusiness struct {
	// 继承
	_ base.BaseBusiness
}

var ImageBusiness *imageBusiness = new(imageBusiness)
