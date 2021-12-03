package producer

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
)

type Producer struct {
	Producer    sarama.SyncProducer
	Topic       string
	ProducerID  int //生产者Id
	PartitionId int32
	MessageId   int
}

func (p *Producer) InitProducer() {
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{define.SERVER_LIST}, config)
	if err != nil {
		tlog.Error("producer closed, err:", err)
		return
	}

	p.Producer = client
	p.Topic = define.TOPIC
	p.ProducerID = ProducerId
	p.PartitionId = PartitionId
	p.MessageId = 1

	ProducerId++
	PartitionId++
}

func (p *Producer) SendMessage() {
	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = p.Topic
	msg.Partition = p.PartitionId
	txt := fmt.Sprintf("ProducerID:%d 我想去的partition:%d this is a test log %d",
		p.ProducerID, p.PartitionId, p.MessageId)
	msg.Value = sarama.StringEncoder(txt)

	// 发送消息
	pid, offset, err := p.Producer.SendMessage(msg)
	//_, _, err := client.SendMessage(msg)
	if err != nil {
		fmt.Println("send msg failed, err:", err)
		return
	}
	tlog.Info(fmt.Sprintf("ProducerID:%d pid:%v offset:%v msg:%s",
		p.ProducerID, pid, offset, txt))

	p.MessageId++
}

func (p *Producer) Close() {
	p.Producer.Close()
}
