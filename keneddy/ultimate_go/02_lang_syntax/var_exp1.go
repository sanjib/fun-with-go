package main

import (
	"fmt"
	"unsafe"
)

func main() {
	var a int
	var b string
	var c float64
	var d bool

	fmt.Printf("a %T %[1]v\n", a)
	fmt.Printf("b %T %[1]v\n", b)
	fmt.Printf("c %T %[1]v\n", c)
	fmt.Printf("d %T %[1]v\n", d)

	fmt.Println("a size", unsafe.Sizeof(a))
	fmt.Println("b size", unsafe.Sizeof(b))
	fmt.Println("c size", unsafe.Sizeof(c))
	fmt.Println("d size", unsafe.Sizeof(d))

	bb := "hello"
	fmt.Println("bb size", unsafe.Sizeof(bb))
	fmt.Println(&bb)

	var i1 int8 = 10
	var i1x = int64(i1)
	var i2 int16 = 10
	var i3 int32 = 10
	var i4 int64 = 10

	fmt.Println("i1 size", unsafe.Sizeof(i1))
	fmt.Println("i1x size", unsafe.Sizeof(i1x))
	fmt.Println("i2 size", unsafe.Sizeof(i2))
	fmt.Println("i3 size", unsafe.Sizeof(i3))
	fmt.Println("i4 size", unsafe.Sizeof(i4))
}
