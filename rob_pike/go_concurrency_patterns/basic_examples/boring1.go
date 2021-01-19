package main

import (
	"fmt"
	"time"
)

func boring1(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Second)
	}
}

func main() {
	boring1("boring!")
}
