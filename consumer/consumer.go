package consumer

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
)

type Consumer struct {
	Consumer   sarama.Consumer
	Topic      string
	ConsumerId int //消费者Id
}

func (c *Consumer) InitConsumer() error {
	consumer, err := sarama.NewConsumer([]string{define.SERVER_LIST}, nil)
	if err != nil {
		return err
	}
	c.Consumer = consumer
	c.Topic = define.TOPIC
	c.ConsumerId = ConsumerId
	ConsumerId++
	return nil
}

//指定partition
//offset 可以指定，传-1为获取最新offest
func (c *Consumer) GetMessage(partitionId int32, offset int64) {
	if offset == -1 {
		offset = sarama.OffsetNewest
	}
	pc, err := c.Consumer.ConsumePartition(c.Topic, partitionId, offset)
	if err != nil {
		tlog.Error("failed to start consumer for partition %d,err:%v", partitionId, err)
		//That topic/partition is already being consumed
		return
	}

	// 异步从每个分区消费信息
	go func(sarama.PartitionConsumer) {
		for msg := range pc.Messages() {
			tlog.Info("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v", c.ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		}
	}(pc)
}

//遍历所有分区
func (c *Consumer) GetMessageToAll(offset int64) {
	partitionList, err := c.Consumer.Partitions(c.Topic) // 根据topic取到所有的分区
	if err != nil {
		tlog.Error("fail to get list of partition:err%v", err)
		return
	}
	tlog.Info("所有partition:", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		c.GetMessage(int32(partition), offset)
	}
}
