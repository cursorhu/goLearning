package main

import (
	"fmt"
	"time"
)

func main() {
	//创建管道，但是不写入值
	ch := make(chan int)
	//创建超时标志的管道，默认值false
	isTimeout := make(chan bool)

	//记录超时
	go func(){
		//休眠10s, 如果休眠期间切到其他goroutine，下次切回来又重头休眠10s
		time.Sleep(10*time.Second)
		//超时：这个goroutine完整休眠了10s !
		isTimeout <- true
	}()

	//主函数监控管道数据
	select {
		case data := <-ch:
			fmt.Println(data)
		case <-isTimeout:
			fmt.Println("isTimeout!")
		//这里不用default，使select在没case满足时能阻塞
		//default:
	}

}