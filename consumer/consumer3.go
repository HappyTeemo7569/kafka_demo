package consumer

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
	"time"
)

type Consumer3 struct {
	Consumer    sarama.Consumer
	Topic       string
	ConsumerId  int //消费者Id
	PartitionId int32
	Offset      int64
}

func (c *Consumer3) InitConsumer() error {
	consumer, err := sarama.NewConsumer([]string{define.SERVER_LIST}, nil)
	if err != nil {
		return err
	}
	c.Consumer = consumer
	c.Topic = define.TOPIC
	c.ConsumerId = ConsumerId
	c.PartitionId = PartitionId
	c.Offset = 0 //先都从头开始 理论上要存在redis

	ConsumerId++
	PartitionId++
	return nil
}

//指定partition
//自己保存partition 和 offset
func (c *Consumer3) GetMessage() {

	pc, err := c.Consumer.ConsumePartition(c.Topic, c.PartitionId, c.Offset)
	if err != nil {
		tlog.Error("failed to start consumer for partition %d,err:%v", c.PartitionId, err)
		//That topic/partition is already being consumed
		return
	}

	// 异步从每个分区消费信息
	//go func(sarama.PartitionConsumer) {
	for msg := range pc.Messages() {
		tlog.Info("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v", c.ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		c.Offset = msg.Offset
	}
	//}(pc)
}

//遍历所有分区
func (c *Consumer3) GetMessageToAll() {
	partitionList, err := c.Consumer.Partitions(c.Topic) // 根据topic取到所有的分区
	if err != nil {
		tlog.Error("fail to get list of partition:err%v", err)
		return
	}
	tlog.Info("所有partition:", partitionList)

	//for partition := range partitionList { // 遍历所有的分区
	for {
		c.GetMessage()
		time.Sleep(1 * time.Second)
	}

	//}
}
