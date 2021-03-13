package mod01

import (
	"fmt"
	"strings"
)

func FizzBuzz(n int) {
	var out []string
	for i := 1; i <= n; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			out = append(out, "Fizz Buzz")
		case i%3 == 0:
			out = append(out, "Fizz")
		case i%5 == 0:
			out = append(out, "Buzz")
		default:
			out = append(out, fmt.Sprintf("%d", i))
		}
	}
	if len(out) > 0 {
		fmt.Println(strings.Join(out, ", "))
	}
}
