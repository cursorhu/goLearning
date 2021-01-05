package main

import "fmt"

//两种不同类型的入参，无返回值
func foo(a int, b string) {
	fmt.Println(a, b)
}

//两种同类型的入参，有返回值(返回列表是类型)
func add(a, b int) int {
	ret := a + b
	return ret
}

//多返回值(返回列表是类型)
func returnValueByType(a, b int) (int, int) {
	return a, b
}

//多返回值，返回局部变量
//注意返回的变量，不能和入参同名(不能定义成 a, b int)，否则重定义报错
func returnValueByVar(a, b int) (ret, err int) {
	//这里是赋值而不是声明赋值，因为函数的返回变量列表已经声明了
	ret = a + b
	err = 0
	//已经声明了要返回的变量，直接return即可
	return
}


func main() {
	//无返回值函数
	foo(1, "hello")
	
	//单返回值函数
	ret := add(1, 2)
	fmt.Println(ret)

	//多返回要用多个变量接收，不想收的可以用'_'省略
	a, b := returnValueByType(1, 2)
	fmt.Println(a, b)
	
	//ret, err是多返回值函数的常见用法，既返回值又返回错误码
	ret, err := returnValueByVar(1, 2)
	fmt.Println(ret, err)
}
