package consumer

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
)

//传入消费者数量
func Get(consumerNum int) {

	offest := int64(0)

	//直接创建多个消费者
	for i := 0; i < consumerNum; i++ {
		consumer := new(Consumer)
		err := consumer.InitConsumer()
		if err != nil {
			tlog.Error("fail to init consumer, err:%v", err)
			return
		}
		consumer.GetMessageToAll(offest)
	}
	//这样会带来重复消费

	//下面试试多个partition

}
