package consumer

import (
	"github.com/HappyTeemo7569/teemoKit/tlog"
	"sync"
)

var nextOffset = int64(0)

//添加一个取号器
var myLock sync.Mutex

func getOffset(consumerId int) int64 {
	tlog.Debug("想要 获取offset:上锁:", consumerId, nextOffset)
	myLock.Lock()
	tlog.Debug("获取offset:上锁:", consumerId, nextOffset)
	return nextOffset
}

func setOffset(consumerId int) {
	nextOffset++
	tlog.Debug("设置offset:解锁:", consumerId, nextOffset)
	myLock.Unlock()
}
