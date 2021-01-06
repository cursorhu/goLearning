package main

import (
	"fmt"
)

func testPanicDefer() {
	fmt.Println("Enter testPanicDefer")
	defer fmt.Println("Exit testPanicDefer")
	panic("panic occur")
}

func main() {
	testPanicDefer()
	fmt.Println("Exit main")
}

