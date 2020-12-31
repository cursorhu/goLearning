package main

import (
	"fmt"
	"sync"
)

func main() {

	//声明sync.Map，不能用常规的make创建
	var scene sync.Map

	//用Store, 初始化将键值对保存到sync.Map
	scene.Store("greece", 97)
	scene.Store("london", 100)
	scene.Store("egypt", 200)

	//用Load, 根据键取到值, 并返回是否有该键
	fmt.Println(scene.Load("london"))

	//用Delete, 根据键删除对应的键值对
	scene.Delete("london")

	//用 Range(回调函数) 形式，遍历所有sync.Map中的键值对
	//匿名func实现回调，Range函数会将遍历的键值写入该回调函数的参数k, v中
	//k, v参数是interface{}类型的，能适配各种类型的map键值类型
	//返回bool类型，若为true继续遍历，false停止遍历
	scene.Range(func(k, v interface{}) bool {
		fmt.Println("iterate:", k, v)
		return true
	})

}
