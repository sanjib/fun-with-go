package main

import (
	"fmt"
	"reflect"
)

func fnPtrExp1(a *int) {
	fmt.Println("a", a, reflect.TypeOf(a))
	fmt.Println("*a", *a, reflect.TypeOf(*a))

	// changes a in main
	//*a = 5

	// changes a in main
	//b := 5
	//*a = b

	// does not change a in main
	b := 5
	a = &b
	fmt.Println("a add in fn", a)
}

func main() {
	a := 1
	fmt.Println(a)
	fnPtrExp1(&a)
	fmt.Println(a)
	fmt.Println("a add in main", &a)
}
