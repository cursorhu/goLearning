package main

import (
	"fmt"
	"net/http"
	"io/ioutil"
	"sync"
)

func main() {

	//声明一个等待组变量
	var wg sync.WaitGroup

	//准备一系列的网站地址
	var urls = []string{
		"https://www.google.com/",
		"https://www.kernel.org/",
		"http://www.github.com/",
	}

	//遍历这些地址
	for _, url := range urls {

		//每创建一个任务，将等待组增加1
		wg.Add(1)

		//创建一个并发任务，处理http访问和接收
		go func(url string) {

			//使用defer，任务完成时将等待组值减1
			defer wg.Done()
			
			//使用http.Get访问指定的地址，返回数据是resp
			resp, err := http.Get(url)
			if err != nil {
		        fmt.Println("http.Get err=", err)
		        return
		    }
		    //resp结构体数据解析成字节数据
		    data, err := ioutil.ReadAll(resp.Body)
		    if err != nil {
		        fmt.Println("ioutil.ReadAll err=", err)
		        return
		    }

			//显示url和返回的数据
			fmt.Println(url, string(data))

		}(url)
	}

	//主任务等待所有创建的任务完成
	wg.Wait()

	fmt.Println("done")
}
