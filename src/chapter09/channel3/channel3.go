package main

import (
	"fmt"
)

func main() {
	
	ch := make(chan interface{}, 3)
	fmt.Println("ch length:", len(ch))
	
	ch <- "hello 1"
	ch <- "hello 2"
	ch <- "hello 3"
	fmt.Println("ch length:", len(ch))

	//带缓存管道可以在一个goroutine内读写
	//for-range可以读管道数据
	//管道读完数据必须手动break，无数据可读会阻塞，可能死锁
	for data := range ch {
		fmt.Println(data)
		if len(ch) == 0 {
			break
		}
	}
	fmt.Println("ch length:", len(ch))

}