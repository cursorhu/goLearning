package main
//导入单个包，fmt包用于格式化输出功能
import "fmt"

func main(){
	//Println: print + line next，打印并换行
	//Println首字母大写，表示是全局公开的方法，对fmt以外的包可见，可调用
	fmt.Println("hello world!");
}