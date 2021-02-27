package main

import (
	"fmt"
	"unsafe"
)

func iter(n int) []struct{} {
	return make([]struct{}, n)
}

func iter2(n int) []int {
	return make([]int, n)
}

func main() {
	for i := range iter(7) {
		fmt.Println(i)
	}
	fmt.Println("--")
	for i := range iter2(7) {
		fmt.Println(i)
	}
	fmt.Println("--")
	fmt.Println("iter 7 size", unsafe.Sizeof(iter(7)))
	fmt.Println("iter2 7 size", unsafe.Sizeof(iter2(7)))
}
