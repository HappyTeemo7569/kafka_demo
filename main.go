package main

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"kafkaDemo/producer"
)

func main() {
	tlog.Info("开始")

	producer.Put()

}
