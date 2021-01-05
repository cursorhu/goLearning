package main

import "fmt"

func main() {

	for x := 0; x < 10; x++ {
		for y := 0; y < 10; y++ {
			if y == 2 {
				// 跳转到标签
				goto Exit
			}
		}
	}
	//不会执行到这里
	return

Exit:
	fmt.Println("done")
}
