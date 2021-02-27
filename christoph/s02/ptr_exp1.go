package main

import (
	"fmt"
)

func main() {
	var a int = 1
	var p *int
	fmt.Println(p)
	//fmt.Println(*p) // will cause panic
	p = &a
	fmt.Println(p)
	fmt.Println(*p)
}
