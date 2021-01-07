package main

import (
	"fmt"
)

// 定义一个数据写入器, interface{}类型表示任意类型
type DataWriter interface {
	WriteData(data interface{}) error
	CanWrite() bool
}

// 定义文件结构，用于实现DataWriter
type file struct {
}

// 实现DataWriter接口的WriteData方法
// 注意入参类型，函数名，返回类型和接口定义完全一致
func (d *file) WriteData(data interface{}) error {
	// 模拟写入数据
	fmt.Println("WriteData:", data)
	return nil
}

func main() {

	// 实例化file
	f := new(file)

	// 声明一个DataWriter的接口
	var writer DataWriter

	// 注册接口的实现：将实现接口的对象(*file类型)，赋值给接口对象(DataWriter类型)
	writer = f

	// 使用DataWriter接口的WriteData方法，实际执行的是*file指向的WriteData
	writer.WriteData("data")
}
