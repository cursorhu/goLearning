package main
import "fmt"

func main() {
	//切片, string, map
	arr := []int{1,2,3,4}
	str := "hello"
	m := map[string]int {
		"up": 8,
		"down": 2,
	}

	for k, v := range arr {
		fmt.Println("key:", k, "value:", v)
	}
	for k, v := range str {
		//这里用Println会直接输出ASCII值而非字符
		fmt.Printf("key: %d, value: %c\n", k, v)
	}
	for k, v := range m {
		fmt.Println("key:", k, "value:", v)
	}

	//管道
	c := make(chan int)
	//匿名函数写法
	go func() {
		c <- 1
		c <- 2
		//关闭管道
		close(c)
	}()

	//管道是在goroutine执行，可能比for-range后执行
	//如果管道还没数据，for-range阻塞
	//管道来一个数据，这里得到一个数据
	//最后管道close, goroutine结束，for-range才会结束
	for v := range c {
		fmt.Println("value:", v)
	}
}