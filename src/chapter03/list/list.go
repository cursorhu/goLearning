package main

import (
	"container/list"
	"fmt"
)

func main() {
	//初始化链表
	l := list.New()

	//查看New出来的list是什么类型(*list.List)
	fmt.Printf("%T\n", l)

	//尾部添加
	l.PushBack("Go")
	//头部添加
	l.PushFront(666)
	//PushBack和PushFront都返回插入成员的指针：＊list.Element
	e := l.PushBack("&")

	//配合InsertAfter、InsertBefore, 在指定成员前后，再插入成员
	l.InsertBefore("Rock", e)
	l.InsertAfter("Roll", e)

	traverselist(l)

	//配合Remove，删除指定成员
	l.Remove(e)

	traverselist(l)
}

//遍历链表
func traverselist(l *list.List) {
	//注意list的起始判断，和成员迭代方式
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
