package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println("int", rand.Int())
	fmt.Println("intn", rand.Intn(10))
	fmt.Println("f32", rand.Float32())
	fmt.Println("f64", rand.Float64())

	sBefore := []string{"apple", "pear", "grape", "orange", "kiwi", "melon"}
	var sAfter []string
	indexes := rand.Perm(len(sBefore))
	fmt.Println("perm", indexes)
	for i := 0; i < len(indexes); i++ {
		sAfter = append(sAfter, sBefore[indexes[i]])
	}
	fmt.Println("s before", sBefore)
	fmt.Println("s after", sAfter)

	const a, b = 10, 50
	for i := 0; i < 10; i++ {
		fmt.Printf("%d ", a+rand.Intn(b-a+1))
	}
	fmt.Println("")

	// random uppercase character
	const c, d = 65, 90
	for i := 0; i < 10; i++ {
		fmt.Printf("%c ", c+rand.Intn(d-c+1))
	}
	fmt.Println("")

	// random uppercase character
	for i := 0; i < 10; i++ {
		fmt.Printf("%c ", 'A'+rune(rand.Intn('Z'-'A'+1)))
	}
	fmt.Println("")
}
