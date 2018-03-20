package Id

import (
	"fmt"
	_ "time"
)

type IdChannel struct {
	processNum int
	request    chan int8
	response   chan int64
}

/**
 * 构造方法.
 *
 * @params int processNum 启动进程数.
 *
 * @return IdChannel
 */
func (this *IdChannel) New(processNum int) *IdChannel {
	this = new(IdChannel)
	this.processNum = processNum
	this.request = make(chan int8, 10000)
	this.response = make(chan int64, 10000)

	return this
}

/**
 * 启动Id生成器.
 */
func (this *IdChannel) Start() {
	for i := 0; i < this.processNum; i++ {
		go func(processId int) {
			fmt.Println("启动Id生成器", processId)

			var idCreator *IdCreator
			idCreator = new(IdCreator)
			params := make(map[string]interface{})

			for true {
				_, close := <-this.request
				if !close {
					fmt.Println("退出Id生成器", processId)
					break
				}
				this.response <- idCreator.GetId(params)

			}
		}(i)
	}
}

/**
 * 停止Id生成器.
 */
func (this *IdChannel) Stop() {
	close(this.request)
}

/**
 * 获取Id.
 */
func (this *IdChannel) GetId() int64 {
	this.request <- 1 // 发出请求
	res := <-this.response
	return res
}
