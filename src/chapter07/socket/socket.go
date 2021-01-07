package main

import (
	"fmt"
	"io"
)

//io包中声明的Writer和Closer接口：
/*
type Writer interface {
	Write(p []byte) (n int, err error)
}

type Closer interface {
	Close() error
}
*/

//定义Socket结构体和方法，实现io包的接口
type Socket struct {

}

func (s *Socket) Write(p []byte) (n int, err error) {
	fmt.Println("Socket.Write")
	return 0, nil
}

func (s *Socket) Close() error {
	fmt.Println("Socket.Close")
	return nil
}


//使用io.Writer的代码，其并不知道Socket和io.Closer的存在
func doWriter(writer io.Writer) {
	fmt.Println("do io.Writer.Write")
	writer.Write(nil)
}

//使用io.Closer的代码，其并不知道Socket和io.Writer的存在
func doCloser(closer io.Closer) {
	fmt.Println("do io.Closer.Close")
	closer.Close()
}


func main() {
	//实例化Socket对象
	s := new(Socket)

	//这里通过传参，注册接口的实现
	//隐式的将Socket类型转换成io.Writer和io.Closer类型，完成接口实现类型的注册
	doWriter(s)
	doCloser(s)
}