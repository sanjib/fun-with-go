package main

import (
	"fmt"
	"unsafe"
)

type example1 struct {
	flag    bool
	counter int16
	pi      float32
}

type example2 struct {
	flag    bool
	counter int32
	pi      float64
}

var example3 struct {
	flag    bool
	counter int16
	pi      float32
}

func main() {
	fmt.Printf("%#v\n", example1{})
	fmt.Println(unsafe.Sizeof(example1{}))

	var e1 example1
	fmt.Printf("%#v\n", e1)
	fmt.Println(unsafe.Sizeof(e1))

	fmt.Printf("%#v\n", example2{})
	fmt.Println(unsafe.Sizeof(example2{}))

	fmt.Println(unsafe.Sizeof(struct {
		flag1 bool
		flag2 bool
	}{}))

	fmt.Println(unsafe.Sizeof(struct {
		pi      float32
		counter int16
		flag    bool
	}{}))

	fmt.Println(unsafe.Sizeof(struct {
		flag    bool
		counter int32
		pi      float32
	}{}))

	fmt.Println(unsafe.Sizeof(struct {
		flag    bool
		counter int64
		pi      float32
	}{}))

	fmt.Println(unsafe.Sizeof(struct {
		counter int64
		pi      float32
		flag    bool
	}{}))

	fmt.Printf("%#v\n", example3)

	type bill struct {
		flag    bool
		counter int32
		pi      float32
	}

	type alice struct {
		flag    bool
		counter int32
		pi      float32
	}

	var b bill
	var a alice
	b = bill(a)
	fmt.Println(a, b)
}
