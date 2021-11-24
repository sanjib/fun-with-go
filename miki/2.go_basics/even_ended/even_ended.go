package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {
	start := time.Now()
	var count int

	min := 1000
	max := 9999

	for i := min; i <= max; i++ {
		for j := i; j <= max; j++ {
			n := i * j
			s := strconv.Itoa(n)
			if s[0] == s[len(s)-1] {
				count++
			}

		}
	}

	d := time.Since(start)
	fmt.Println("found", count, "even-ended numbers")
	fmt.Println("time took", d.Seconds(), "secs.")
}
