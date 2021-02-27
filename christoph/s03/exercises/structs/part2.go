package main

import (
	"fmt"
	"unsafe"
)

type emptyStruct struct{}
type notEmptyStruct struct {
	x int
}
type metaEmptyStruct struct {
	A struct{}
	B struct{}
}

func main() {
	es := emptyStruct{}
	fmt.Println("size emptyStruct", unsafe.Sizeof(es))

	nes := notEmptyStruct{}
	fmt.Println("size notEmptyStruct", unsafe.Sizeof(nes))

	mes := metaEmptyStruct{
		A: struct{}{},
		B: struct{}{},
	}
	fmt.Println("size metaEmptyStruct", unsafe.Sizeof(mes))

	// side test - slice with a million len and cap is only 24 bytes
	sliceWithLargeCapLen := make([]int, 1e6, 1e6)
	fmt.Println("sliceWithLargeCapLen", sliceWithLargeCapLen)
	fmt.Println("sliceWithLargeCapLen", unsafe.Sizeof(sliceWithLargeCapLen))

}
