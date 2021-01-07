package main
import "fmt"

//定义一个抽象的myType类
type myType int

//实现其方法
func (m *myType) IsZero() bool{
	return *m == 0
}

func (m *myType) Add(v myType) myType {
	//修改了对象
	*m += v
	return *m
}

func main() {
	m := new(myType)
	//调用方法
	fmt.Println(m.IsZero())
	//10强转成myType类型
	fmt.Println(m.Add(myType(10)))
	fmt.Println(m.IsZero())
}