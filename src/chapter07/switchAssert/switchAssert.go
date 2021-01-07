package main
import "fmt"

func assertType(v interface{}) {
	switch v.(type) {
		case int:
			fmt.Printf("%v, %T\n", v, v)
		case string:
			fmt.Printf("%v, %T\n", v, v)
		case map[string]string:
			fmt.Printf("%v, %T\n", v, v)
		case struct {
				a int
				s string
			}:
			fmt.Printf("%v, %T\n", v, v)
	}
}

func main() {
	assertType(1024)
	assertType("hello")
	assertType(map[string]string {
		"W": "up",
		"S": "down",
	})
	assertType(struct {
		a int
		s string
	} {
		a: 100,
		s: "hundred",
	})
}