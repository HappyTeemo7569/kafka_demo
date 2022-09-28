package kafka

import (
	"log"
	"testing"
)

func Test_Get(t *testing.T) {

	topic := "test_log"
	var kafkaConsumer = KafkaConsumer{
		Node:  []string{Conn},
		Topic: topic,
	}

	kafkaConsumer.MessageQueue = make(chan []byte, 1000)
	go kafkaConsumer.Consume()

	for {
		msg := <-kafkaConsumer.MessageQueue
		deal(msg)
	}

}

func deal(msg []byte) {
	log.Printf(string(msg))
}
