package consumer

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
)

func Get(ConsumerId int) {
	consumer, err := sarama.NewConsumer([]string{define.SERVER_LIST}, nil)
	if err != nil {
		tlog.Error("fail to start consumer, err:%v\n", err)
		return
	}
	partitionList, err := consumer.Partitions(define.TOPIC) // 根据topic取到所有的分区
	if err != nil {
		tlog.Error("fail to get list of partition:err%v\n", err)
		return
	}
	fmt.Println(string(partitionList))

	//partition := 0
	//sarama.OffsetNewest

	for partition := range partitionList { // 遍历所有的分区
		// 针对每个分区创建一个对应的分区消费者
		pc, err := consumer.ConsumePartition(define.TOPIC, int32(partition), 1)
		if err != nil {
			tlog.Error("failed to start consumer for partition %d,err:%v\n", partition, err)
			return
		}
		defer pc.AsyncClose()
		// 异步从每个分区消费信息
		go func(sarama.PartitionConsumer) {
			for msg := range pc.Messages() {
				tlog.Info("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v \n", ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
			}
		}(pc)
	}
}
