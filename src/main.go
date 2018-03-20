package main

import (
	"fmt"
	"id"
	"time"
)

func main() {
	var exitSignal chan int16
	exitSignal = make(chan int16, 1)

	fmt.Println("Hello World")

	var httpServerIdCreator *Id.Http.HttpWorker
	httpServerIdCreator.Start()

	var idCreator *Id.IdCreator
	var id int64
	params := make(map[string]interface{})

	idCreator = new(Id.IdCreator)

	id = idCreator.GetId(params)
	fmt.Println(id)
	params["result"] = int64(0xFF)
	id = idCreator.GetId(params)
	fmt.Println(id)
	delete(params, "result")
	id = idCreator.GetId(params)
	fmt.Println(id)

	/** 多协程直接调用测试 平均每秒5800w条 单进程每微妙58条
	var ts, te, count int64
	ts = time.Now().Unix()
	for true {
		id = idCreator.GetId(params)
		// fmt.Println(id)
		count++
		if count > 1000000000 {
			te = time.Now().Unix()
			fmt.Println(te - ts)
			break
		}
	}
	*/

	/** 多协程直接调用测试 平均每秒1333w条 单进程每微妙14条
	var ts, te, count int64
	for i := 0; i < 16; i++ {
		go func(pid int) {
			ts = time.Now().Unix()
			for true {
				id = idCreator.GetId(params)
				// fmt.Println(id)
				count++

				if count > 1000000000 {
					te = time.Now().Unix()
					fmt.Println(pid, te-ts)
					break
				}
			}
		}(i)
	}
	/*/

	var coreIdCreator *Id.IdChannel
	coreIdCreator = coreIdCreator.New(16)
	coreIdCreator.Start()

	id = coreIdCreator.GetId()
	fmt.Println(id)

	go func() {
		fmt.Println("Time Start")
		time.Sleep(5 * time.Second)
		fmt.Println("Time End")
		exitSignal <- 0
	}()

	/** 单进程管道调用效率测试 平均每秒227w条
	var ts, te, count int64
	ts = time.Now().Unix()
	for true {
		id = coreIdCreator.GetId()
		// fmt.Println(id)
		count++
		if count > 100000000 {
			te = time.Now().Unix()
			fmt.Println(te - ts)
			break
		}
	}
	*/

	/** 多进程同时获取id 平均每秒222w条
	var ts, te, count int64
	for i := 0; i < 16; i++ {
		go func(pid int) {
			ts = time.Now().Unix()
			for true {
				id = coreIdCreator.GetId()
				// fmt.Println(id)
				count++

				if count > 100000000 {
					te = time.Now().Unix()
					fmt.Println(pid, te-ts)
					break
				}
			}
		}(i)
	}
	*/

	// 持续运行
	e := <-exitSignal
	fmt.Println("Exit with code", e)
	coreIdCreator.Stop()
	time.Sleep(5 * time.Second)

}
