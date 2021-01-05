package main
import "fmt"

func main() {
	//匿名函数声明后直接调用
	//直接在函数大括号后接小括号，传入参数
	func(a int) {
		fmt.Println(a)
	} (100)

	//匿名函数赋值给函数变量
	f := func(a int) {
		fmt.Println(a)
	}
	//调用
	f(100)

	//匿名函数作为回调函数
	//此处以匿名函数的形式，传入了具体方法：打印
	//此处是直接在调用函数时，定义传入的变量
	visit([]int{1,2,3,4}, func(v int) {
		fmt.Println(v)
	})
}

//visit:以某种方式, 访问切片数组的每个成员
//具体方式取决于传入的函数
func visit(arr []int, f func(int)) {
	for _, v := range arr {
		f(v)
	}
}