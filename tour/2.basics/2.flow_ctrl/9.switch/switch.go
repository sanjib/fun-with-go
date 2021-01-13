package main

import (
	"fmt"
	"runtime"
	"strings"
)

func main() {
	fmt.Print("Go is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux.")
	default:
		fmt.Printf("%s.\n", strings.Title(os))
	}
}
