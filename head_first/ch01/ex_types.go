package main

import (
	"fmt"
	"reflect"
)

func main() {
	// fmt.Println(reflect.TypeOf(42))
	// fmt.Println(reflect.TypeOf('A'))
	// fmt.Println(reflect.TypeOf("A"))

	fmt.Println(reflect.TypeOf(25))
	fmt.Println(reflect.TypeOf(true))
	fmt.Println(reflect.TypeOf(5.2))
	fmt.Println(reflect.TypeOf(1))
	fmt.Println(reflect.TypeOf(false))
	fmt.Println(reflect.TypeOf(1.0))
	fmt.Println(reflect.TypeOf("hello"))

}
