package main
import "fmt"

func foo() {
	fmt.Println("foo")
}

func main() {
	//声明一个函数变量
	var f func()
	//赋值
	f = foo
	//调用
	f()
}