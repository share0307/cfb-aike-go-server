package queue

import (
	"aike-cfb-server/kernel/helper"
	"fmt"
	"github.com/go-redis/redis"
)

type HandleReceiveMsgProcess func(string)

/**
	通用实现
*/
type CommonQueueImplementation struct {
	// 队列名称
	queueName string
	// 是否开启去重标记
	isDuplicateCheckFlag bool
	// 重复检测的map
	duplicateMap	map[string]string
	// redis组建
	rds *redis.Client
	// x 秒之内的任务不得重复
	duplicateLifeCycle int

	HandleReceiveMsgProcess HandleReceiveMsgProcess
}

/**
	设置队列名称
 */
func (c *CommonQueueImplementation)SetQueueName(name string)  {
	c.queueName = name

	c.HandleReceiveMsgProcess("aaa")
}

/**
	获取队列名称
 */
func (c *CommonQueueImplementation)GetQueueName() string {
	return c.queueName
}


/**
是否开启去重判断
*/
func (c *CommonQueueImplementation)SetEnableDuplicateCheckFlag(flag bool) {
	c.isDuplicateCheckFlag = flag
}

/**
	是否开启去重判断
 */
func (c *CommonQueueImplementation)IsEnableDuplicateCheck() bool {
	return c.isDuplicateCheckFlag
}

/**
	设置去重map
 */
func (c *CommonQueueImplementation)SetDuplicateMap(duplicateMap map[string]string) {
	c.duplicateMap = duplicateMap
}

/**
	生成 sign 的规则
	todo：此处有个坑，哪怕在子类中覆盖重写的，因为接收者的不同，所以也不能算是覆盖，所以无法重写！！
*/
func (c *CommonQueueImplementation)GetDuplicateMap() map[string]string {
	return c.duplicateMap
}

/**
	去重的规则
	当map为空时，则生成去重规则将出错！！
 */
func (c *CommonQueueImplementation)GetDuplicateSign() string {
	DuplicateMap := c.GetDuplicateMap()

	fmt.Println(DuplicateMap)

	if len(DuplicateMap) == 0 {
		panic("请先设置完成 duplicateMap 的设置！！")
	}

	return helper.GetDuplicateSignByMapDefaultSep(DuplicateMap)
}

/**
	设置 x 秒之内，任务不得重复
 */
func (c *CommonQueueImplementation)SetDuplicateLifeCycle(second int) {
	c.duplicateLifeCycle = second
}
/**
	设置 x 秒之内，任务不得重复，设置默认值
 */
func (c *CommonQueueImplementation)SetDuplicateLifeCycleDefault() {
	c.SetDuplicateLifeCycle(300)
}

/**
	默认的去重规则，返回一个空的map，但map为空时，则生成去重规则将出错！！
	X 秒之内不得重复，也就是说重复 key 的生命周期
 */
func (c *CommonQueueImplementation)GetDuplicateLifeCycle() int {
	return c.duplicateLifeCycle
}

/**
	设置默认的redis链接
 */
func (c *CommonQueueImplementation)SetDefaultRds() {
	//c.SetRds()
}

/**
	设置默认的redis链接
 */
func (c *CommonQueueImplementation)SetRds(rds *redis.Client) {
	c.rds = rds
}

// 获取重复判断的中间件，依赖redis
func (c *CommonQueueImplementation)GetRds() *redis.Client {
	// 当 rds 为 nil 时，则设置默认的redis链接，并且返回
	if c.rds == nil {
		c.SetDefaultRds()
	}

	return c.rds
}

// 定义一些通用依赖组建。。。
