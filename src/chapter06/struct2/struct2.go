package main

import (
	"fmt"
)

//People类型的字段是指向该类型的指针，实现了链表
type People struct {
	Name string
	Child *People
}

//嵌套的实例化，键值对方式
//relation是package main内的全局变量
//全局变量的声明只能用var方式，:=只能在函数内使用
var relation = &People {
	Name: "grandpa",
	Child: &People { 
		Name: "dad",
		Child: &People {
			Name: "me",
			//指针未初始化，是默认的nil
		},
	},
}

//表示地址的结构体
type Address struct {
	City string
	Phone string
	Zipcode int
}
//构造函数
func newAddress(c, p string, z int) *Address {
	//直接值实例化，省略键
	return &Address {
		c,
		p,
		z,
	}
}

func main() {
	//%+v只能打印一层结构体字段
	fmt.Printf("%+v\n %+v\n %+v\n", relation, relation.Child, relation.Child.Child)
	
	addr := newAddress("武汉", "027-", 430000)
	fmt.Printf("%+v\n", addr)
}


