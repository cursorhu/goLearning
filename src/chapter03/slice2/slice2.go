package main
import "fmt"

func main(){
	//声明切片变量，没有分配内存
	var strSlice []string
	var intSlice []int
	
	//定义一个空切片，尽管没有初始化值，但已分配了内存
	intSliceEmpty := []int{}
	//定义一个初始化的切片
	intSliceInitialed := []int{1,2,3}
	
	//切片判空
	fmt.Println(strSlice == nil)
	fmt.Println(intSlice == nil)
	fmt.Println(intSliceEmpty == nil)

	//获取切片长度（已有成员的总数）
	fmt.Println(len(strSlice), len(intSlice), len(intSliceEmpty), len(intSliceInitialed))
	//切片变量的类型
	fmt.Printf("%T, %T, %T\n", strSlice, intSlice, intSliceEmpty)
}