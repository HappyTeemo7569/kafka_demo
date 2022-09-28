package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
	"time"
)

var Conn string
var SyncProducer sarama.SyncProducer

func newSyncProducer() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	// 随机的分区类型：返回一个分区器，该分区器每次选择一个随机分区
	config.Producer.Partitioner = sarama.NewRandomPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{Conn}, config)
	//defer producer.Close()

	if err != nil {
		return nil, err
	}

	return producer, nil
}

//随机发送
func Send(topic, content string) (bool, error) {
	if SyncProducer == nil {
		var err error
		SyncProducer, err = newSyncProducer()
		if err != nil {
			return false, nil
		}
	}

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Key:   sarama.StringEncoder(time.Now().String()), //
		Value: sarama.StringEncoder(content),
	}

	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := SyncProducer.SendMessage(msg)
	if err != nil {
		return false, err
	}
	fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	return true, nil
}

var SyncProducerOrderly sarama.SyncProducer

func newSyncProduceOrderly() (sarama.SyncProducer, error) {
	config := sarama.NewConfig()
	// 等待服务器所有副本都保存成功后的响应
	config.Producer.RequiredAcks = sarama.WaitForAll
	//固定分区
	config.Producer.Partitioner = sarama.NewHashPartitioner
	// 是否等待成功和失败后的响应
	config.Producer.Return.Successes = true

	// 使用给定代理地址和配置创建一个同步生产者
	producer, err := sarama.NewSyncProducer([]string{Conn}, config)
	//defer producer.Close()

	if err != nil {
		return nil, err
	}

	return producer, nil
}

//顺序发送
//根据key去hash
func SendOrderly(topic, content, key string) error {
	if SyncProducerOrderly == nil {
		var err error
		SyncProducerOrderly, err = newSyncProduceOrderly()
		if err != nil {
			return nil
		}
	}

	//构建发送的消息，
	msg := &sarama.ProducerMessage{
		Topic: topic,
		//Partition: 1,                         //顺序发送
		Key:   sarama.StringEncoder(key), //根据key去hash
		Value: sarama.StringEncoder(content),
	}

	//SendMessage：该方法是生产者生产给定的消息
	//生产成功的时候返回该消息的分区和所在的偏移量
	//生产失败的时候返回error
	partition, offset, err := SyncProducerOrderly.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Partition = %d, offset=%d\n", partition, offset)
	return nil
}
