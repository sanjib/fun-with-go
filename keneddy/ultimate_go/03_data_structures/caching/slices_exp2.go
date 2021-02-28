package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var data []string
	//data := []string{}
	//dataAlt := make([]string, 5)
	//fmt.Printf("%#v\n", data)
	//fmt.Printf("%#v\n", dataAlt)
	lastCap := cap(data)
	fmt.Printf("%#v, %d\n", data, lastCap)

	var emptyStruct struct{}
	fmt.Printf("%#v, %d\n", emptyStruct, unsafe.Sizeof(emptyStruct))

}
