package producer

import (
	"fmt"
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"github.com/Shopify/sarama"
	"kafkaDemo/define"
)

func Put() {
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
	defer client.Close()

	index := 1
	for {
		// 构造一个消息
		msg := &sarama.ProducerMessage{}
		msg.Topic = define.TOPIC
		txt := fmt.Sprintf("this is a test log %d", index)
		msg.Value = sarama.StringEncoder(txt)

		// 发送消息
		pid, offset, err := client.SendMessage(msg)
		//_, _, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		tlog.Info("pid:%v offset:%v msg:%s \n", pid, offset, txt)

		index++
	}
}
