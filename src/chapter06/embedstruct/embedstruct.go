package main

import "fmt"

/** 常规写法，在结构体内定义变量，没用内嵌**/
//基础颜色:三原色
type BasicColor struct {
	R, G, B float32
}
//完整颜色，包括三原色和透明度
type Color struct {
	//将基本颜色作为成员
	Basic BasicColor
	//透明度
	Alpha float32
}

func ColorInit() {
	var c Color
	c.Basic.R = 255
	c.Basic.G = 255
	c.Basic.B = 255
	c.Alpha = 50
	fmt.Printf("%+v\n", c)
}


/** 推荐写法，在结构体内用匿名类型，内嵌方式**/
type BasicColorEmbed struct {
	R, G, B float32
}

type FullColor struct {
	//将其他结构体类型作为匿名成员
	BasicColorEmbed
	//透明度
	Alpha float32
}

func FullColorInit() {
	var c FullColor
	//方式一：可以对内嵌类型的成员直接赋值
	c.R = 255
	c.G = 255
	c.B = 255
	c.Alpha = 50
	fmt.Printf("%+v\n", c)
	
	//方式二：内嵌类型也可以作为显式的成员
	c.BasicColorEmbed.R = 0
	c.BasicColorEmbed.G = 0
	c.BasicColorEmbed.B = 0
	c.Alpha = 50
	fmt.Printf("%+v\n", c)
}

func main() {
	ColorInit()
	FullColorInit()
}