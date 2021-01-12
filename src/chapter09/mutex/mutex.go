package main
import (
	"fmt"
	"time"
	"sync"
)

// 定义能安全计数的结构体，包括变量和锁
type SafeCounter struct {
	cnt int
	Guard sync.Mutex
}

func (c *SafeCounter) Get() int {

	// 调用sync.Mutex的Lock方法
	c.Guard.Lock()
	// 延迟调用sync.Mutex的UnLock方法
	defer c.Guard.Unlock()
	// 以下语句属于加锁的区域，不会有数据竞态
	return c.cnt
}

func (c *SafeCounter) Add(val int) {

	c.Guard.Lock()
	defer c.Guard.Unlock()
	c.cnt += val
}

func main() {
	//创建一个SafeCounter的实例
	counter := &SafeCounter{}
	//创建10个协程，每个都对共享的main进程的counter对象+1
	for i := 0; i < 10; i++ {
		go func() {
			counter.Add(1)
		}()
	}
	
	//main进程休眠等待一下
	time.Sleep(time.Second)
	//打印结果
	fmt.Println(counter.Get())
}