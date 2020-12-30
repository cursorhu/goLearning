package main

import(
	"fmt"
	"math/rand"
	"time"
)

//数据生产者
func producer(header string, channel chan string){
	//无限循环，不停地生产数据
	for{
		//将随机数和字符串格式化为字符串发送给通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		//等待 1 秒, 在goroutine中执行时，不影响其他goroutine的执行
		time.Sleep(time.Second)
	}
}

//数据消费者
func consumer(channel chan string){
	//不停地获取数据
	for{
		//从通道中取出数据 ，如果没有，此处会阻塞直到信道中返回数据
		message := <- channel
		//打印数据
		fmt.Println(message)
	}
}

func main(){
	//创建一个字符串类型的通道
	channel := make(chan string)
	//创建两个并发goroutine(go协程)，执行producer函数
	go producer("cat", channel)
	go producer("dog", channel)
	//在主线程(main所在线程)，执行consumer函数
	consumer(channel)
}