package main

import (
	"fmt"
	"math/rand"
	"time"
)

func boring2(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		duration := time.Duration(rand.Intn(1e3))
		fmt.Println("duration", duration)
		time.Sleep(duration * time.Millisecond)
	}
}

func main() {
	boring2("boring!")
}
