package base

import "errors"

/**
	枚举基类
 */
type BaseConstant struct {
	
}

/**
	映射枚举值的地方
	todo：必须要重写此方法
 */
func (c *BaseConstant)GetNames() map[string]interface{} {
	return map[string]interface{}{}
}

/**
	获取全部的枚举值
 */
func (c *BaseConstant)All() []string {
	keys := make([]string, 0, len(c.GetNames()))

	for k,_ := range c.GetNames(){
		keys = append(keys, k)
	}

	return keys
}

/**
	获取每个枚举的值
 */
func (c *BaseConstant)GetName(code string) (interface{}, error) {
	if c.Has(code) {
		return c.GetNames()[code], nil
	}

	return "",errors.New("值不存在！")
}

/**
	返回某个枚举值是否合法
 */
func (c *BaseConstant)Has(code string) bool {
	if _,exists := c.GetNames()[code]; exists {
		return true
	}

	return false
}
