package main

import (
	"fmt"
)

//累加器
func Accumulate(value int) func() int {

	//返回一个匿名函数，其中value是外部环境：Accumulate函数的变量，因此返回的是闭包
	return func() int {
		value++
		return value
	}
}

func main() {

	//创建一个累加器，返回的是闭包
	//闭包是在main创建的，其匿名函数和变量(value)的生命周期和main相同
	accumulator := Accumulate(0)

	//对闭包中变量累加
	fmt.Println(accumulator())
	fmt.Println(accumulator())

	//打印累加器的地址
	fmt.Printf("%p\n", accumulator)

	//又创建一个累加器闭包
	accumulator2 := Accumulate(1)
	fmt.Println(accumulator2())
	fmt.Printf("%p\n", accumulator2)

	//直接创建和使用闭包
	fmt.Printf("%p\n", Accumulate(2))
}
