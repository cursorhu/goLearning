package main

import (
	"fmt"
)

type man struct {
	name string
	age int
}

func main() {
	me := man{"cursorhu", 26}
	p := &me
	fmt.Println(p.name, p.age)
}