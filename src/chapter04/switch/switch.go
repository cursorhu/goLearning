package main
import "fmt"

func main() {
	//常规用法，判断变量，单分支自动break
	var a = "hello"
	switch a {
		case "hello":
			fmt.Println(1)
		case "world":
			fmt.Println(2)
		default:
			fmt.Println(0)
	}

	//多个case合并
	switch a {
		case "hello", "world":
			fmt.Println(1)
		default:
			fmt.Println(0)
	}

	//判断表达式, switch后无变量, 在case内判断表达式
	switch {
		case a == "hello": 
			fmt.Println(1)
		default:
			fmt.Println(0)
	}

}
