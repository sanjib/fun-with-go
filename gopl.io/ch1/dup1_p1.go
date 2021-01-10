// My practice version of dup1. I prefer the vars:
//	lines (vs counts)
//	scanner (vs input)
// more in my version.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	lines := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		lines[scanner.Text()]++
	}
	
	for line, n := range lines {
		fmt.Printf("%d\t%s\n", n, line)
	}
	fmt.Println(lines)
}
