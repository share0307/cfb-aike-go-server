package queue

import (
	"context"
	"sync"
)

/**
	实例化mq服务
 */
func NewMqService() *mqService {
	return &mqService{

	}
}

/**
	提供队列服务
 */
type mqService struct {
	ctx context.Context
	queueList []IQueue
}

/**
	添加队列服务
 */
func (mq *mqService)Add(queue IQueue)  {
	mq.queueList = append(mq.queueList, queue)
}

/**
	运行队列
 */
func (mq *mqService)Run(ctx context.Context, group *sync.WaitGroup) {
	for i := 0;i < len(mq.queueList); i++ {
		go func(idx int) {
			// 调用队列本身的Init方法，初始化服务等组件
			mq.queueList[idx].InitConsumer(ctx, group)

			// 进消息监听，并且把消息推送给固定方法
			consumeChan, err := mq.queueList[idx].Consume()

			if err != nil {
				panic(err)
			}

			go func() {
				for delivery := range consumeChan {
					mq.queueList[idx].HandleReceiveMsgProcess(&delivery)
				}
			}()
		}(i)
	}

}