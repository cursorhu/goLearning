package main

import (
	"fmt"
	//这个包可以处理命令行参数
	//更多参考：https://www.jianshu.com/p/4ec4f674bda3
	"flag" 
)

func main() {
	//flag.String的三个参数： 名称，默认值，注释说明
	//返回命令行参数对应的值的指针(此时未初始化)
	var param = flag.String("ops", "", "operations")
	//解析命令行参数，写入param指针
	flag.Parse()

	//map的声明定义
	//分解一下： map[string]func()是声明一个map: 其key是string, value是func()匿名函数
	//大括号内是定义map的键值，{"key":value,} 格式。其中value是匿名函数
	m := map[string]func(){
		"startup": func(){
			fmt.Println("ops: startup")
		},
		"shutdown": func(){
			fmt.Println("ops: shutdown")
		},
	}

	//获取map的值，即匿名函数，赋值给变量f
	if f, ok := m[*param]; ok {
		f()
	} else {
		fmt.Println("ops not found")
	}

}