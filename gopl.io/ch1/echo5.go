// Modify the echo program to print the index and value of each of
// its arguments, one per line.
package main

import (
	"fmt"
	"os"
)

func main() {
	s, sep := "", ""
	for index, arg := range os.Args[1:] {
		s += fmt.Sprintf("%s%d.%s", sep, index+1, arg)
		sep = "\n"
	}
	fmt.Println(s)
}
