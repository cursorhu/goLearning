package main

import (
	"fmt"
	"time"
	"sync/atomic"
)

// 全局的序列号
var (
	seq int64
)

// 原子操作的序列号累加器
func GenID() int64 {
	/*错误写法：atomic修改的seq和return seq还是有竞态*/
	//atomic.AddInt64(&seq, 1)
	//return seq

	/*正确写法：直接return atomic的返回值*/
	return atomic.AddInt64(&seq, 1)
}

func main() {

	// 10个并发任务，生成序列号（return的数据被丢弃）
	for i := 0; i < 10; i++ {
		go GenID()
	}
	//让main等一下goroutine
	time.Sleep(time.Second)
	// 返回最终结果
	// 这里不能直接返回seq，否则竞态
	//fmt.Println(seq)
	fmt.Println(GenID())
}
