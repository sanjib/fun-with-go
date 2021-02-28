package main

import (
	"fmt"
	"unsafe"
)

func modSlice(s []string) {
	//s[4] = "Plum"
	s = append(s, "Plum")
}

func main() {
	fruits := make([]string, 5)
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"

	modSlice(fruits)

	fmt.Println("fruits", fruits, unsafe.Sizeof(fruits), len(fruits), cap(fruits))

	//fruits[5] = "Cherry"
	//fruits = append(fruits, "Cherry")
	//fmt.Println("fruits", fruits, unsafe.Sizeof(fruits), len(fruits), cap(fruits))

}
