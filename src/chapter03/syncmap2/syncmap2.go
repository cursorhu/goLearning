package main

import (
	"fmt"
	"sync"
	"time"
)

func main(){
	var msync sync.Map
	go mapWriter(&msync)
	go mapReader(&msync)

	for{} //主函数等待
}

//注意要传指针，不能是sync.Map形参
func mapWriter(m *sync.Map) {
	for {
		//写键值对<"A", 1>到sync.Map
		m.Store("A", 1)
		//sleep用于goroutine切换
		time.Sleep(time.Second)
	}
}

func mapReader(m *sync.Map) {
	for{
		//读"A"对应的值
		fmt.Println(m.Load("A"))
		time.Sleep(time.Second)
	}
}