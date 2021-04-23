package queue

import (
	"aike-cfb-server/kernel/helper"
	"aike-cfb-server/provider"
	"context"
	"fmt"
	"github.com/go-redis/redis"
	"github.com/streadway/amqp"
	"sync"
)

/**
	通用实现
*/
type CommonQueueImplementation struct {
	// 队列配置别名
	queueConfigAlias string
	// 是否开启去重标记
	isDuplicateCheckFlag bool
	// 重复检测的map
	duplicateMap	map[string]string
	// redis组建
	rds *redis.Client
	// x 秒之内的任务不得重复
	duplicateLifeCycle int

	// mq的服务提供者
	mqProvider *provider.RabbitmqProvider

	// 当进程退出时，需要做的一些析构方法
	downFunc func()
}

/**
	监控队列
 */
func (c *CommonQueueImplementation)Monitor(ctx context.Context, group *sync.WaitGroup) {
	// 添加分组任务
	group.Add(1)

	// 正式监听
	go func() {
		// 进行监听 context
		select {
		case <- ctx.Done():
			{
				// 关闭通用的方法
			}

			fmt.Println("退出了！！")

			// 执行关闭方法
			if c.downFunc != nil {
				c.downFunc()
			}

			// 完成分组任务
			group.Done()
		}
	}()
}

/**
	设置关闭时执行的方法
 */
func (c *CommonQueueImplementation)SetDownFunc(downFunc func()) {
	c.downFunc = downFunc
}

/**
	设置队列
 */
func (c *CommonQueueImplementation)SetQueueConfig(alias string)  {
	c.queueConfigAlias = alias
}

/**
	链接队列
 */
func (c *CommonQueueImplementation)InitMq()  {
	// 获取实例
	mqProvider := provider.NewRabbitmqProvider(c.queueConfigAlias)

	// 链接
	mqProvider.Connect()

	// 绑定交换机与队列的关系
	mqProvider.InitExchangeAndQueue()

	c.mqProvider = mqProvider
}

/**
	消费
 */
func (c *CommonQueueImplementation)Consume() (<- chan amqp.Delivery, error) {
	consumeChan, err := c.mqProvider.Consume()

	if err != nil {
		fmt.Println(err)
	}

	return consumeChan, err
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

func (c *CommonQueueImplementation)PublishSimpleMsg(body []byte) {
	publishing := amqp.Publishing{}

	publishing.Body = body
	publishing.ContentType = "text/plain"

	c.PublishMsg(publishing)
}

// mq操作的一些方法
/**
	生产数据
 */
func (c *CommonQueueImplementation)PublishMsg(publishing amqp.Publishing) {
	err := c.mqProvider.Publish(publishing)

	fmt.Println(err)
}
