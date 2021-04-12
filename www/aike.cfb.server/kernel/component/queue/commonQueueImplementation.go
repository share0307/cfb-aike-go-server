package queue

import (
	"aike-cfb-server/kernel/helper"
	"fmt"
	"github.com/go-redis/redis"
)

/**
	通用实现
*/
type CommonQueueImplementation struct {
	Rds *redis.Client
}

/**
消息出错的流程通用实现
todo：默认不作任何业务
*/
func (q *CommonQueueImplementation)HandleMsgErrProcess()  {
	return
}

/**
	是否开启去重判断
 */
func (q *CommonQueueImplementation)IsEnableDuplicateCheck() bool {
	return false;
}


/**
生成 sign 的规则
todo：此处有个坑，哪怕在子类中覆盖重写的，因为接收者的不同，所以也不能算是覆盖，所以无法重写！！
*/
func (h *CommonQueueImplementation)GetDuplicateMap() map[string]string {
	return map[string]string{}
}

/**
	去重的规则
	当map为空时，则生成去重规则将出错！！
 */
func (q *CommonQueueImplementation)GetDuplicateSign() string {
	duplicateMap := q.GetDuplicateMap()

	fmt.Println(duplicateMap)

	if len(duplicateMap) == 0 {
		panic("请先设置完成 GetDuplicateMap() 的设置！！")
	}

	// 先排序
	return helper.GetDuplicateSignByMapDefaultSep(duplicateMap)
}

/**
	返回去重的组合
 */
//func (q *CommonQueueImplementation)GetDuplicateMap()  map[string]string{
//	return map[string]string{}
//}

/**
	默认的去重规则，返回一个空的map，但map为空时，则生成去重规则将出错！！
 */
// X 秒之内不得重复，也就是说重复 key 的生命周期
func (q *CommonQueueImplementation)GetDuplicateLifeCycle() int {
	return 300
}

// 设置重复判断的中间件，依赖redis
//func (q *CommonQueueImplementation)SetDuplicateRds(rds *redis.Client) {
//
//}

/**
	设置默认的redis链接
 */
func (q *CommonQueueImplementation)SetDefaultDuplicateRds() {
	//q.Rds =
}

// 获取重复判断的中间件，依赖redis
func (q *CommonQueueImplementation)GetDuplicateRds() *redis.Client {
	// 当 Rds 为 nil 时，则设置默认的redis链接，并且返回
	if q.Rds == nil {
		q.SetDefaultDuplicateRds()
	}

	return q.Rds;
}

// 定义一些依赖组建。。。
