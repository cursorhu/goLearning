package main

import (
	"fmt"
)

func main() {
	
	ch := make(chan interface{})

	go func(){
		ch <- "hello 1"
		ch <- "hello 2"
		ch <- "hello 3"
		ch <- "hello 4"
	}()

	//方式一：读出数据，并赋值变量
	data := <- ch
	fmt.Println(data)

	//方式二：读出并一次性使用数据
	fmt.Println(<-ch)

	//方式三：非阻塞方式读出
	data, ok := <-ch
	if ok {
		fmt.Println(data)
	}

	//方式四：读出并丢弃数据
	<-ch

}
