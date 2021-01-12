package main

import (
	"fmt"
	"net/http"
)

func main() {

	//声明一个channel，用于同步
	c := make(chan int, 3)
	

	//准备一系列的网站地址
	var urls = []string{
		"https://www.google.com/",
		"https://www.kernel.org/",
		"http://www.github.com/",
	}

	//遍历地址, 创建任务
	for _, url := range urls {

		//创建一个并发任务，处理http访问和接收
		go func(url string) {

			//使用http.Get访问指定的地址，丢弃返回数据
			_, err := http.Get(url)
			if err != nil {
		        fmt.Println("http.Get err=", err)
		        //return //这里不能return，否则管道没写入
		    }
		    
			//显示url
			fmt.Println(url)

			//每完成一个任务，向channel写入数据
			//注意：管道读写，不能用defer语句
			c <- 1

		}(url)
	}

	// 等待N个任务完成，注意这里用cap而不是len，cap能表示总共等几个任务
	// 在没有任务完成时，channel为空，此处会阻塞。 若有任务完成，完成一个取出一个。
    for i := 0; i < cap(c); i++ {
        <-c
    }

	fmt.Println("done")
}