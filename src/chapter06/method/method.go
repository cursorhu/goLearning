package main
import "fmt"

//定义Bag结构体类型，相当于C++ class类
type Bag struct {
	items []int
}
//声明Bag类型的方法：Insert, Bag的指针作为接收器，相当于C++ this对象指针
func (b *Bag) Insert(id int) {
	b.items = append(b.items, id)
}

func main() {
    //分配一个Bag类型的对象
	b := new(Bag)
	//调用方法
	b.Insert(100)
	fmt.Println(b.items[0])
}

