package main

import (
	sanlib "../../../sanjib/lib"
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(sanlib.Seed())
	c := make(chan string)
	go boring4("boring!", c)

	for i := 0; i < 5; i++ {
		fmt.Printf("You say: %q\n", <-c)
	}
	fmt.Println("You're boring. I'm leaving.")
}

func boring4(msg string, c chan string) {
	for i := 0; ; i++ {
		d := time.Duration(rand.Intn(1e3)) * time.Millisecond

		c <- fmt.Sprintf("%s %d %v", msg, i, d)
		time.Sleep(d)
	}
}
