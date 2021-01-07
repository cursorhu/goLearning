package main

import (
	"fmt"
)

//定义服务接口Service
type Service interface {
	Start()
	Log(string)
}

//log模块
type Logger struct {

}

//log模块的方法, 实现Service的Log()
func (l *Logger) Log(s string) {
	fmt.Println(s)
}

//服务启动模块
type Starter struct {

}

//实现Service的Start()
func (s *Starter) Start() {
	fmt.Println("Starter.Start")
}

//定义实现完整服务接口的结构体，包含各子模块，实现完整接口
type Servicer struct {
	Starter
	Logger
}

func main() {
	//实现方赋值给接口方，完成接口注册
	var s Service = new(Servicer)
	//使用接口
	s.Start()
	s.Log("log success")
}