package main

import (
	"fmt"
	"runtime"
)

// 自定义的，panic时要传递的上下文信息
type panicContext struct {
	errinfo string // 所在函数
	errcode int
}

// 保护方式允许一个函数
func ProtectRun(entry func()) {

	//延迟处理的函数
	defer func() {

		//发生宕机时，recover()能获取panic()函数内的信息，返回给err
		err := recover()

		switch err.(type) {
		case runtime.Error: // 运行时错误
			fmt.Println("runtime error:", err)
		default: // 非运行时错误，打印err值
			fmt.Println("error:", err)
		}

	}()

	entry()

}

func main() {
	fmt.Println("运行前")

	//测试手动触发的错误处理
	ProtectRun( func() {

		fmt.Println("手动宕机前")

		// 使用panic传递上下文
		panic( &panicContext{
			"手动触发panic",
			400,
		} )

		fmt.Println("手动宕机后")
	})

	// 测试运行时错误：故意造成空指针访问错误
	ProtectRun(func() {

		fmt.Println("赋值宕机前")

		var a *int
		*a = 1

		fmt.Println("赋值宕机后")
	})

	fmt.Println("运行后")
}
