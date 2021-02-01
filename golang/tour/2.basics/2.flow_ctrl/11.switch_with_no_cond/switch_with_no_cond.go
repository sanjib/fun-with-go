package main

import (
	"fmt"
	"time"
)

func main() {
	hour := time.Now().Hour()
	switch {
	case hour < 12:
		fmt.Println("Good morning.")
	case hour < 17:
		fmt.Println("Good afternoon.")
	default:
		fmt.Println("Good evening.")
	}
}
