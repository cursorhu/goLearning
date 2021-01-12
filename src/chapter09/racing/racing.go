package main

import "fmt"

func addNumber(i int) int {
	//创建goroutine做计算
	go func() {
		i++
	}()
	//此处return的i是goroutine执行之前还是之后？
	//不确定，即对数据i的访问上，产生竞态
	return i
}

func main() {
	var i int
	fmt.Println(addNumber(i))
}

