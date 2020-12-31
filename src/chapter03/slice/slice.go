package main

import(
	"fmt"
)
//从数组生成切片
func main(){
	var arr[10] int

	//注意这里不能用range，因为数组还没初始化
	for i := 0; i < 10; i++ {
		arr[i] = i
	} 
	//左闭右开，从下标5，到第10个成员的切片(不包含第10个成员)
	fmt.Println(arr[5:10])
	//从下标5，到最后成员，的切片
	fmt.Println(arr[5:])
	//所有成员的切片
	fmt.Println(arr[:])
	//所有成员清除，结果是没成员，而不是成员为0
	fmt.Println(arr[0:0])
}