package main

import (
	"fmt"
	"time"
)

func main() {
	go sayHello()
	time.Sleep(1)
}

func sayHello() {
	fmt.Println("hello")
}
