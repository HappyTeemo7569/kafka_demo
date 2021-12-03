package consumer

import "github.com/HappyTeemo7569/teemoKit/tlog"

//传入消费者数量
func Get(consumerNum int) {

	/*
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
	*/

	/*
		//试试取号器
		consumer := new(Consumer2)
		err := consumer.InitConsumer()
		if err != nil {
			tlog.Error("fail to init consumer, err:%v", err)
			return
		}
		consumer.GetMessageToAll()
		//提供的是获取的消息，没有单独获取一条消息
	*/

	//下面试试多个partition
	for i := 0; i < consumerNum; i++ {
		consumer := new(Consumer3)
		err := consumer.InitConsumer()
		if err != nil {
			tlog.Error("fail to init consumer, err:%v", err)
			return
		}
		go consumer.GetMessageToAll()
	}

}
