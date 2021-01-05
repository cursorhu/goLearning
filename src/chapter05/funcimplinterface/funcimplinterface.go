package main

import (
	"fmt"
)

// 声明接口
type Invoker interface {
	// 声明Call方法，需要其他地方实现
	Call(interface{})
}

// 定义结构体类型
type Struct struct {
}

// 用结构体，实现Call方法
func (s *Struct) Call(p interface{}) {
	fmt.Println("from struct", p)
}

// 函数定义为类型
type FuncCaller func(interface{})

// 用函数，实现Call方法
func (f FuncCaller) Call(p interface{}) {
	// 调用f函数本体
	f(p)
}

func main() {

	// 声明接口变量
	var invoker Invoker

	// 实例化结构体
	s := new(Struct)

	// 将实例化的结构体赋值到接口，实现接口和方法的关联
	invoker = s

	// 使用接口调用实例化结构体的方法Struct.Call
	invoker.Call("hello")

	// 将匿名函数转为FuncCaller类型，再赋值给接口，实现接口和方法的关联
	invoker = FuncCaller(func(v interface{}) {
		fmt.Println("from function", v)
	})

	// 使用接口调用FuncCaller.Call，内部会调用函数本体
	invoker.Call("hello")
}
