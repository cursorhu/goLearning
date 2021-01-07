package main

import (
	"fmt"
)

func main() {

	/*1.空接口的赋值*/
	var any interface{}
	any = 100
	fmt.Printf("%v, %T\n", any, any)
	any = "hello"
	fmt.Printf("%v, %T\n", any, any)

	/*空接口向其他变量赋值*/
	//不能直接赋值
	//var s string = any //编译报错
	
	//应该使用断言，做接口的类型转换
	s, ok := any.(string)
	if ok {
		fmt.Printf("%v, %T\n", s, s)
	}

	//接口相等，是值相等，类型相同
	var a interface{} = 0x41
	var b interface{} = string(0x41)
	fmt.Println(a == b, a, b)

	//不能把动态类型的数据赋值给接口，然后比较
	//动态类型包括：切片，map
	//可以比较的是：结构体，函数变量，同一个通道，定长数组
	
	//以下示例导致崩溃，因为切片是动态的类型
	/*
	var c interface{} = []int{10}
	var d interface{} = []int{5}
	fmt.Println(c == d) //崩溃
	*/
}