package main

import (
	"fmt"
	"runtime"
)

func main() {
	numCPU := runtime.NumCPU()
	fmt.Println(numCPU)
}
