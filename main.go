package main

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"kafkaDemo/consumer"
	"kafkaDemo/producer"
	"time"
)

func main() {
	tlog.Info("开始")

	go producer.Put()
	go consumer.Get(3)

	for {
		time.Sleep(time.Hour * 60)
	}
}
