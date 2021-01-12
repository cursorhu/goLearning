package main

import (
	"fmt"
)

func main() {
	
	ch := make(chan interface{})

	/* 错误写法：同一个goroutine内使用channel
	ch <- "hello"
	fmt.Println(<-ch)
	*/

	/* 没生效的写法：在goroutine内打印，不会显示
	go func(){
		fmt.Println(<-ch)
	}()
	*/

	/* 不必要写法： 把channel变量当普通变量传参
	go func(c chan interface{}){
		c <- "hello"
	}(ch)
	*/

	//推荐写法：goroutine做数据运算和修改，channel变量直接用
	go func(){
		ch <- "hello"
	}()

	//main routine处理结果，此处自带阻塞
	fmt.Println(<-ch)
}

