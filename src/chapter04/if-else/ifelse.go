package main
import "fmt"

func main(){
	 //常用写法
	 //注意else不能换行写，否则报错
    var a int = 11
    if a > 10 {
        fmt.Println("a > 10")
    } else {
        fmt.Println("a <= 10")
    }
    //以下是另一种写法，变量声明和作用域都在if-else，可以避免重复定义
    if a := 11; a >10 {
        fmt.Println("a > 10")
    } else {
        fmt.Println("a <= 10")
    }

}