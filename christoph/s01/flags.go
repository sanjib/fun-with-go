package main

import (
	"flag"
	"fmt"
)

func main() {
	verbose := flag.Bool("v", false, "Verbose")
	count := flag.Int("c", 1, "Count")

	flag.Parse()

	fmt.Printf("%#v\n", *verbose)
	fmt.Printf("%#v\n", *count)

	fmt.Println(flag.Args())
}
