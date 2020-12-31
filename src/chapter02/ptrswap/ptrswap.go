//ptrswap.go
package main
import(
	"fmt"
)

func main(){
	a, b := 1, 2
	pa := &a
	pb := &b
	//注意，带格式化输出只能Printf, Println会原样输出字符串内的 %d, %p字符
	fmt.Printf("Before swapval: a: %d, b: %d, pa: %p, pb: %p \n", a, b, pa, pb)
	swapval(pa, pb)
	fmt.Printf("After swapval: a: %d, b: %d, pa: %p, pb: %p \n", a, b, pa, pb)
	swapptr(pa, pb)
	fmt.Printf("After swapptr: a: %d, b: %d, pa: %p, pb: %p \n", a, b, pa, pb)
	pb = &a
	fmt.Printf("After change pb: pa: %p, pb: %p \n", pa, pb)
}

func swapval(a, b *int){
	//这里把指针指向的变量改了
	t := *a
	*a = *b
	*b = t
}

func swapptr(a, b *int){
	//这里交换的是形参，实参不受影响
	a, b = b, a
}