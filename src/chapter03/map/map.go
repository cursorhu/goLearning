package main
import (
	"fmt"
	"sort"
)

func main(){
	//make方式创建map
	m := make(map[string]int)
	//初始化一个键值对成员
	m["A"] = 1;
	//查找一个键对应的值
	fmt.Println(m["A"])
	//若键不存在，值为默认值，即0
	fmt.Println(m["Z"])
	//更好的查找，能判断键值对是否存在
	v, ok := m["Z"]
	fmt.Println(v, ok)

	//直接方式创建map,大括号形式，类似JSON格式
	keyboardmap := map[string]string{
		"W": "up",
		"A": "left",
		"S": "down",
		"D": "right",
	}

	//遍历已初始化的map
	for k, v := range keyboardmap {
		fmt.Println(k, v)
	}

	//只获取值，或只获取键的遍历
	for _, v := range keyboardmap {
		fmt.Println(v)
	}
	//这里有两种写法，k,_ := 或直接 k :=
	for k, _ := range keyboardmap {
		fmt.Println(k)
	}

	//map是无序的，若需要按键排序输出，要放到切片再手动排序
	//声明一个切片数组，用来存储map的key
	arr := []string{}
	//遍历map, 写入key到切片
	for k, _ := range keyboardmap {
		arr = append(arr, k)
	}
	//对切片的key成员排序
	sort.Strings(arr)
	//有序输出map的<k, v>
	//注意遍历切片，range返回是 下标,成员值
	for _, k := range arr {
		fmt.Println(k, keyboardmap[k])
	}

	//删除键值对
	delete(keyboardmap, "W")
	for k, v := range keyboardmap {
		fmt.Println(k, v)
	}

	//清空map
	/*没有清空map的方法，等着垃圾回收吧 ^_^! */
}