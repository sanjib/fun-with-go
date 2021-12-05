package main

import (
	"fmt"
	"log"
	"strings"
)

func main() {
	var sb strings.Builder

	sb.WriteString("This is a string 1\n")
	sb.WriteString("This is a string 2\n")
	sb.WriteString("This is a string 3\n")

	log.Println(sb.String())

	fmt.Println("capacity:", sb.Cap())
	sb.Grow(1024)
	fmt.Println("capacity:", sb.Cap())

	for i := 0; i <= 10; i++ {
		fmt.Fprint(&sb, "string %d -- ", i)
	}
	fmt.Println(sb.String())
	fmt.Println("builder size:", sb.Len())

}
