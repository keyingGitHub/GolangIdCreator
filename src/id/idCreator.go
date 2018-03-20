package Id

import (
	_ "fmt"
	"time"
)

const TIME_L uint = 52
const SERVER_L uint = 2
const PROCESS_L uint = 4
const ORDERNUM_L uint = 6

type IdCreator struct {
	orderNum  int64
	timestamp int64
	lastts    int64
	serverNo  int64
	processNo int64
}

/**
 * 生成一个新的可用Id.
 */
func (this *IdCreator) GetId(params map[string]interface{}) int64 {
	var result int64

	ts := this.getTimeStamp()
	ts = ts << (ORDERNUM_L + SERVER_L + PROCESS_L)

	orderNum := this.getOrderNum()
	orderNum = orderNum << (SERVER_L + PROCESS_L)

	serverId := this.getServerId()
	serverId = serverId << PROCESS_L

	processId := this.getProcessId()

	result = ts | orderNum | serverId | processId

	return result
}

/**
 * 获取机器Id.
 */
func (this *IdCreator) getServerId() int64 {
	return 0
}

/**
 * 获取进程Id.
 */
func (this *IdCreator) getProcessId() int64 {
	return 0
}

/**
 * 生成微秒时间戳.
 */
func (this *IdCreator) getTimeStamp() int64 {
	this.timestamp = time.Now().UnixNano() / 1000
	return this.timestamp
}

/**
 * 生成时间内序号.
 */
func (this *IdCreator) getOrderNum() int64 {
	if this.timestamp != this.lastts {
		// 新时间
		this.lastts = this.timestamp
		this.orderNum = 0
	}

	orderNum := this.orderNum
	this.orderNum++
	return orderNum
}

/*
32 10 10
2 4
6
*/
