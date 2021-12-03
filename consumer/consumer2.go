package consumer

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
)

type Consumer2 struct {
	Consumer   sarama.Consumer
	Topic      string
	ConsumerId int //消费者Id
}

func (c *Consumer2) InitConsumer() error {
	consumerConfig := sarama.NewConfig()
	consumerConfig.Version = sarama.V2_8_0_0
	consumerConfig.Consumer.Return.Errors = false
	consumerConfig.Consumer.Offsets.AutoCommit.Enable = false // 禁用自动提交，改为手动
	consumerConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

	consumer, err := sarama.NewConsumer([]string{define.SERVER_LIST}, consumerConfig)
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
func (c *Consumer2) GetMessage(partitionId int32) {

	for {
		offest := getOffset(c.ConsumerId) //取号码
		pc, err := c.Consumer.ConsumePartition(c.Topic, partitionId, offest)
		if err != nil {
			tlog.Error("failed to start consumer for partition %d,err:%v", partitionId, err)
			return
		}

		// 异步从每个分区消费信息
		// 读一条就释放

		//但是这里是直接取了所有的。
		for msg := range pc.Messages() {
			tlog.Info("ConsumerId:%d Partition:%d Offset:%d Key:%v Value:%v",
				c.ConsumerId, msg.Partition, msg.Offset, msg.Key, string(msg.Value))

			setOffset(c.ConsumerId)
			//sarama.Com
			break
		}
	}

}

//遍历所有分区
func (c *Consumer2) GetMessageToAll() {

	partitionList, err := c.Consumer.Partitions(c.Topic) // 根据topic取到所有的分区
	if err != nil {
		tlog.Error("fail to get list of partition:err%v", err)
		return
	}
	tlog.Info("所有partition:", partitionList)

	for partition := range partitionList { // 遍历所有的分区
		c.GetMessage(int32(partition))
	}
}
