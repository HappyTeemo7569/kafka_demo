package kafka

import (
	"fmt"
	"github.com/Shopify/sarama"
)

//防止宕机，可以自己存一下offest

func getOffsetCacheKey(topic string, partitionId int32) string {
	return fmt.Sprintf("kafka:offset:%s:%d", topic, partitionId)
}

func store(key string, offset int64) {
	//伪代码，存redis
}

func setConsumeOffset(topic string, partitionId int32, offset int64) {
	//利用redis存储
	//store(getOffsetCacheKey(topic, partitionId), offset)
}

func exist(key string) bool {
	//伪代码，判断redis key 是否存在
	return false
}

func get(key string) int64 {
	//伪代码，读redis
	return 0
}

func getConsumeOffset(topic string, partitionId int32) (offset int64) {
	key := getOffsetCacheKey(topic, partitionId)

	if exist(key) {
		return get(key) + 1
	}

	//默认从最新开始
	setConsumeOffset(topic, partitionId, sarama.OffsetNewest)
	return sarama.OffsetNewest
}
