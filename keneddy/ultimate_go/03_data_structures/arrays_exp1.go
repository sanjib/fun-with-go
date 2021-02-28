package main

import "fmt"

func main() {
	var fruits [5]string
	fruits[0] = "Apple"
	fruits[1] = "Orange"
	fruits[2] = "Banana"
	fruits[3] = "Grape"
	fruits[4] = "Plum"
	fmt.Println("fruits", fruits)

	numbers := [4]int{10, 20, 30, 40}
	fmt.Println(numbers)
}
