package main

import "fmt"

//声明结构体，用于测试多种数据类型的传参机制
type Data struct {
	num int //测试普通类型
	p *int //测试指针类型
	arr []int //测试切片
	m map[string]string //测试map
	struc strucType //测试结构体
}

//传入数据中的结构体类型
type strucType struct {
	name string
	age int
}

//指针结构体，用于记录传入函数的数据结构体的各成员的地址是否变化
type ptrStrucType struct {
	ptrInteger *int
	ptrPointer **int
	ptrArr *[]int
	ptrMap *map[string]string
	ptrStruc *strucType	
}

//测试函数
func passByValue(in Data) Data {

	//输出函数内各成员的值
	fmt.Printf("in Func data's value: %+v\n", in)
	//输出函数内各成员的地址，这个ptrgroup是形参，所以用声明定义
	ptrgroup := ptrStrucType{&in.num, &in.p, &in.arr, &in.m, &in.struc}
	fmt.Printf("in Func data's ptrgroup: %+v\n", ptrgroup)
	//函数内数据的起始地址
	fmt.Printf("in Func data's ptr: %p\n", &in)

	return in
}

func main() {

	//定义传入函数的数据
	in := Data{
		
		num: 1,

		p: new(int),

		arr: []int {1, 2, 3},

		m: map[string]string {
			"w": "up",
			"s": "down",
		},

		struc: strucType{
			"cursorhu",
			26,
		},

	}

	ptrgroup := ptrStrucType{&in.num, &in.p, &in.arr, &in.m, &in.struc}

	//输入数据各成员的值, %+v用于打印结构体变量
	fmt.Printf("input data value: %+v\n", in)
	//输入数据各成员的地址
	fmt.Printf("input data's ptrgroup: %+v\n", ptrgroup)
	//输入结构的指针地址
	fmt.Printf("input data's ptr: %p\n", &in)

	//结构体传入函数，返回同类型的结构体
	out := passByValue(in)

	//输出数据各成员的值
	fmt.Printf("output data value: %+v\n", out)
	//输出数据各成员的地址
	ptrgroup = ptrStrucType{&out.num, &out.p, &out.arr, &out.m, &out.struc}
	fmt.Printf("output data's ptrgroup: %+v\n", ptrgroup)
	//输出数据的起始地址
	fmt.Printf("output data's ptr: %p\n", &out)
}
