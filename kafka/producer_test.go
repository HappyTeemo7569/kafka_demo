package kafka

import (
	"fmt"
	"testing"
	"time"
)

func Test_Send(t *testing.T) {
	for {
		index := 0
		go func() {
			_, err := Send("test_log", fmt.Sprintf("lox_%d", index))
			if err != nil {
				t.Log("测试失败:" + err.Error())
				return
			} else {
				t.Log("测试成功")
			}
		}()
		index++

		time.Sleep(1 * time.Second)
	}

}

func Test_newSyncProducer(t *testing.T) {

}

func Test_SendOrderly(t *testing.T) {

	index := 0
	for {

		err := SendOrderly("test_log", fmt.Sprintf("lox_%d", index), "test")
		if err != nil {
			t.Log("测试失败:" + err.Error())
			return
		} else {
			t.Log("测试成功")
		}
		index++
		time.Sleep(1 * time.Second)
	}

}
