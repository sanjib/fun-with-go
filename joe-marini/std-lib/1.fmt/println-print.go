package main

import "fmt"

func main() {
	fmt.Print("Welcome to Go!")
	fmt.Println("Enter the new line!")

	const answer = 42
	fmt.Println("The meaning of life is", answer)

	const a, b, c = 5, 5, 10
	fmt.Println("The sum of", a, "and", b, "and", c, "is", a+b+c)

	items := []int{10, 20, 40, 80}
	n, err := fmt.Println(items)
	fmt.Println("n:", n, "; err:", err)
}
