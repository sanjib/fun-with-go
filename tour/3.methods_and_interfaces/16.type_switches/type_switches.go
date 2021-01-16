package main

import "fmt"

func do(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("int: %v %T\n", v, v)
	case float64:
		fmt.Printf("float64: %v %T\n", v, v)
	case string:
		fmt.Printf("string: %v %T\n", v, v)
	case bool:
		fmt.Printf("bool: %v %T\n", v, v)
	default:
		fmt.Printf("Unknown: %v %T\n", v, v)
	}
}

func main() {
	do(21)
	do(3.14)
	do("hello")
	do(true)
	do([]int{1, 2, 3})

	var x interface{} = "x"
	do(x)
	do(&x)
}
