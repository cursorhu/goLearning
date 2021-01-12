package main
import (
	"fmt"
	"time"
)

func counter() {
	var tick int
	//持续计数
	for {
		tick++
		fmt.Println("tick:", tick)
		//休眠1s
		time.Sleep(time.Second)
	}
}


func main() {
	//创建一个goroutine
	go counter()

	//main函数继续执行，同时counter函数也在并发执行
	//Scanln会接收用户输入，直到回车换行才返回
	//如果输入没完成，Scanln阻塞，此时main和counter并发执行。
	//如果输入完成，Scanln返回，随后main结束。
	//main是进程，进程生命周期结束，内部的goroutine都会被强行结束。
	var input string
	fmt.Scanln(&input)
	fmt.Println(input)
}