package main

import (
	"fmt"
	"io"
)

/*
io包中的WriteCloser是组合的接口, 包含Writer和Closer接口

type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}

type WriteCloser interface {
	Writer
	Closer
}

*/

//声明一个设备结构体
type Device struct {

}

//实现io.Writer和io.Closer的方法
func (d *Device) Write(p []byte) (n int, err error) {
	fmt.Println("Device.Write")
	return 0, nil
}

func (d *Device) Close() error {
	fmt.Println("Device.Close")
	return nil
}

func main() {
	//用实现组合接口的类对象，注册
	var wc io.WriteCloser = new(Device)
	//调用组合接口中子接口的方法
	wc.Write(nil)
	wc.Close()

	//只注册子接口
	var w io.Writer = new(Device)
	w.Write(nil)
	//w.Close() //不能调用，io.Writer没有Close()方法

}