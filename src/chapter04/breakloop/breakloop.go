package main

import "fmt"

func main() {

OuterLoop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(i, j)
				//可以指定跳出哪层循环
				break OuterLoop 
			case 3:
				fmt.Println(i, j)
				break OuterLoop
			}
		}
	}
}
