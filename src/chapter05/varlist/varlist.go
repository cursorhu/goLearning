package main

import (
	"fmt"
	"bytes"
	"os"
)

func joinString(varlist ...string) string {
	//定义一个字节数组缓冲区
	var b bytes.Buffer
	//遍历可变参数列表，得到每个参数
	for _, s := range varlist {
		//参数字符串写入缓冲区
		b.WriteString(s)
	}

	//缓冲区转化成字符串，输出
	return b.String()
}

func main() {
	//直接调用
	fmt.Println(joinString("hello ", "world ", "joinString"))

	//传入其他可变参数
	//os.Args是输入参数的切片数组，...访问其所有成员，相当于转化成了可变参数列表
	//os.Args[i]和C 的argv[i] 一样，首个参数是该程序本身，后面依次是入参
	fmt.Println(joinString(os.Args...))
	//改进一下，传入从第二个到最后一个参数的切片的所有成员
	fmt.Println(joinString(os.Args[1:]...))
}