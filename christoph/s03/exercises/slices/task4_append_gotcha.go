package main

import "fmt"

func main() {
	src := []int{}
	src = append(src, 0)
	src = append(src, 1)
	src = append(src, 2)
	fmt.Printf("src %#v\n", src)
	dest1 := append(src, 3)
	dest2 := append(src, 4)
	fmt.Printf("dest1 %#v\n", dest1)
	fmt.Printf("dest2 %#v\n", dest2)
	fmt.Printf("src %#v\n", src)
}
