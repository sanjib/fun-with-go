package main

import (
	sanlib "../../../sanjib/lib"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(sanlib.Seed())
	go boring3("boring")

	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

func boring3(msg string) {
	for i := 0; ; i++ {
		d := time.Duration(rand.Intn(1e3)) * time.Millisecond
		fmt.Println(msg, i, "next at", d)
		time.Sleep(d)
	}
}
