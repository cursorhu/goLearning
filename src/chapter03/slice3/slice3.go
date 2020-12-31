package main
import "fmt"

func main(){
	//创建切片，成员是int类型，分配3个成员的内存，切片最大容量10个成员
	//如果后续添加成员超过10个，切片会自动扩容，类似C++ vector
	//第三个参数切片容量可省略
	a:= make([]int, 3, 10)
	//打印已分配内存的成员
	fmt.Println(len(a))

	//用append初始化切片
	var arr []int
	for i:=0; i < 10; i++ {
		arr = append(arr, i)
		fmt.Printf("len: %d, cap: %d, arr: %p\n", len(arr), cap(arr), arr)
	}

	//append批量添加成员
	var strArr []string
	strArr = append(strArr, "A", "B", "C")
	fmt.Println(strArr)
	//append批量添加另一个切片的成员
	strArrAppend := []string{"D", "E", "F"}
	strArr = append(strArr, strArrAppend...)
	fmt.Println(strArr)

	//拷贝切片，目标切片先要初始化内存，能放得下源切片的数据
	strArrNew := make([]string, len(strArr))
	//copy格式： copy(dest, src)，返回成功拷贝的成员数
	n := copy(strArrNew, strArr)
	fmt.Println(n)
	fmt.Println(strArrNew)

	//删除成员：目前没有直接方法
	//利用切片区间左闭右开，用先切割，后连接的方法
	index := 1
	strArr = append(strArr[:index], strArr[index+1:]...)
	fmt.Println(strArr)
}